package lib

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type Queue struct {
	Channel      *amqp.Channel
	Queue        amqp.Queue
	ExchangeName string
	Queuekey     string
	QueueName    string
}

type consumerConfig struct {
	QueueName    string
	ConsumerName string
	Handle       consumerHandle
}

//消费函数
type consumerHandle func(string, consumerConfig) int

var (
	QueueConn      *amqp.Connection
	MqStatus       string
	MqHandleSucess int = 1
	MqHandleFail   int = 0
)
var consumers = []consumerConfig{
	consumerConfig{"test_queue", "test_consumer", testHandle},
	consumerConfig{"test_queue1", "test_consumer1", testHandle},
}

func testHandle(s string, c consumerConfig) int {
	logrus.Info(s)
	logrus.Info(c.QueueName)
	if s == "" {
		return MqHandleFail
	}
	return MqHandleSucess //成功是返回1 失败时返回0
}

func init() {
	//url := GetConfig("rabbitmq", "url", "")
	//MqStatus = GetConfig("rabbitmq", "status", "0")
	//if MqStatusOn() {
	//	QueueConn, _ = amqp.Dial(url)
	//	for _, consumer := range consumers {
	//		RunConsumer(consumer)
	//	}
	//}
}

//初始化一个生产者
func InitProducer(exchange string, queueName string, queueKey string) Queue {
	var queue Queue
	queue.Channel, _ = QueueConn.Channel()
	_ = queue.Channel.ExchangeDeclare(exchange, "direct", true, false, false, true, nil)
	queue.Queue, _ = queue.Channel.QueueDeclare(queueName, true, false, false, true, nil)
	_ = queue.Channel.QueueBind(queueName, queueKey, exchange, false, nil)
	queue.QueueName = queueName
	queue.ExchangeName = exchange
	queue.Queuekey = queueKey
	return queue
}

//发送消息
func (q Queue) Send(msg string) {
	_ = q.Channel.Publish(q.ExchangeName, q.Queuekey, true, false, amqp.Publishing{
		Timestamp:    time.Now(),
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(msg),
	})
}

//初始化一个消费者
func InitConsumer(queueName string, consumerName string) <-chan amqp.Delivery {
	var queue Queue
	queue.Channel, _ = QueueConn.Channel()
	consumer, _ := queue.Channel.Consume(queueName, consumerName, false, false, false, true, nil)
	return consumer
}

//启动一个消费者
func RunConsumer(c consumerConfig) {
	consumer := InitConsumer(c.QueueName, c.ConsumerName)
	logrus.Info("队列开启: ", c.ConsumerName)
	go func() {
		for {
			m, ok := <-consumer
			if ok {
				res := c.Handle(string(m.Body), c)
				if res > 0 {
					if err := m.Ack(true); err != nil {
						logrus.Info(err)
					}
				}
			} else {
				time.Sleep(time.Duration(2) * time.Second)
			}
		}
	}()
}

func MqStatusOn() bool {
	return MqStatus == "1"
}

func MqStatusOff() bool {
	return MqStatus == "0"
}
