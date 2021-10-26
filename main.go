package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-api/lib"
	"go-api/route"
	"syscall"
)

func main() {
	gin.SetMode(lib.Server.Model)
	serverPort := fmt.Sprintf("0.0.0.0:%d", lib.Server.Port)
	r := gin.New()
	startRoute(r)
	//r.LoadHTMLGlob("/users/xiashuang/go/src/awesomeProject/src/htmls/**/*")
	server := endless.NewServer(serverPort, r)
	server.BeforeBegin = func(add string) {
		logrus.Printf("pid 是: %d", syscall.Getpid())
	}
	_ = server.ListenAndServe()
}

/**
开启监听路由
*/
func startRoute(r *gin.Engine) {
	route.InitApiRoute(r)
}
