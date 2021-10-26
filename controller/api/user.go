package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": "",
	})

}

func TalentInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": "{}",
	})
}
