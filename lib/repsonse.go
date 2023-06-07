package lib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
}

var (
	Resp = response{}
)

func (r response) SuccessJson(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "成功",
		"data":    data,
	})
}

func (r response) SuccessMsg(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
	})
}

func (r response) ErrorJson(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}
