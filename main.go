package main

import (
	"flag"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-api/lib"
	"go-api/route"
	"syscall"
)

func setParams()  {
	var config_path = flag.String("config_path","","指定配置文件目录")
	var app_path = flag.String("app_path","","指定app目录")
	flag.Parse()
	if *config_path != "" {
		lib.ConfigPath = *config_path
	}
	if *app_path != "" {
		lib.APP_APTH = *app_path
	}
}


func main() {
	setParams()
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
