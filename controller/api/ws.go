package api

import (
	"github.com/gin-gonic/gin"
	"go-api/lib"
	"log"
)

func WsConnect(c *gin.Context) {
	user_id := c.Query("uid")
	log.Println("uid", user_id)
	lib.WsConnection(c, user_id)
}

func TestGet(c *gin.Context) {
	user_id := c.Query("uid")
	log.Println("uid", user_id)
}
