package api

import (
	"github.com/gin-gonic/gin"
	"go-api/lib"
	"net/http"
)

func QueueProducer(c *gin.Context) {
	message := c.Query("msg")
	exchange := c.Query("exchange")
	queue := c.Query("queue")
	queueKey := c.Query("queue_key")
	producer := lib.InitProducer(exchange, queue, queueKey)
	producer.Send(message)
}

func QueueTest(c *gin.Context){
	var k  = 0
	for i := 1; i <= 100000000; i++ {
		k +=  i
	}
	c.JSON(http.StatusOK,k)
}
