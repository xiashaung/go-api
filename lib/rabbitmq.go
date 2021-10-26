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

//消费函数
type consumerHandle func(string) int

var (
	QueueConn *amqp.Connection
)

func init() {
	url := GetConfig("rabbitmq", "url", "")
	if url != "" {
		QueueConn, _ = amqp.Dial("amqp://admin:admin@localhost:5672/")
		RunConsumer("test_queue", "test_consumer", func(s string) int {
			logrus.Info(s)
			return 1 //成功是返回1 失败时返回0
		})
	}
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
	q.Channel.Publish(q.ExchangeName, q.Queuekey, true, false, amqp.Publishing{
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
func RunConsumer(queueName string, consumerName string, callback consumerHandle) {
	consumer := InitConsumer(queueName, consumerName)
	go func() {
		for {
			m, ok := <-consumer
			if ok {
				res := callback(string(m.Body))
				if res > 0 {
					if err := m.Ack(true); err != nil {
						logrus.Info(err)
					}
				}

			} else {
				logrus.Info("channel closed")
				break
			}
		}
	}()
}
