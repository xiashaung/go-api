package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-api/lib"
)

func QueueProducer(c *gin.Context) {
	message := c.Query("msg")
	producer := lib.InitProducer("test_exchange", "test_queue", "test_key")
	producer.Send(message)
}

func QueueTest(c *gin.Context){
	var k  = 0
	for i := 1; i <= 100000000; i++ {
		k +=  i
	}
	c.JSON(http.StatusOK,k)
}
