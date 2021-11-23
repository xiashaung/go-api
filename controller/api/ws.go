package api

import (
	"github.com/gin-gonic/gin"
	"go-api/lib"
)

func WsConnect(c *gin.Context)  {
	user_id := c.GetInt("uid")
	lib.WsConnection(c,user_id)
}
