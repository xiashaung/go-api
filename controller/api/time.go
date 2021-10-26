package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"go-api/lib"
)

func ToTimestamp(c *gin.Context) {
	datetime := c.PostForm("datetime")
	if datetime == "" {
		lib.ErrorJson(c, 10001, "请输入正确时间!")
	}
	timestamp, _ := time.ParseInLocation(lib.Server.TimeFormat, datetime, time.Local)
	lib.SuccessJson(c, gin.H{"timestamp": timestamp.Unix()})
}

func Nowtime(c *gin.Context) {
	lib.SuccessJson(c, gin.H{"time": time.Now().Unix()})
}

//
func ToDatetime(c *gin.Context) {
	timestamp := c.PostForm("timestamp")
	if timestamp == "" {
		lib.ErrorJson(c, 10001, "请输入正确时间!")
	}
	formatTime, _ := strconv.ParseInt(timestamp, 10, 64)
	datetime := time.Unix(formatTime, 0).Format(lib.Server.TimeFormat)
	lib.SuccessJson(c, gin.H{"datetime": datetime})
}
