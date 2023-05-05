package main

import (
	"flag"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go-api/lib"
	"go-api/route"
	"log"
	"syscall"
)

func setParams() {
	var config_path = flag.String("config_path", "", "指定配置文件目录")
	var app_path = flag.String("app_path", "", "指定app目录")
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
	serverPort := fmt.Sprintf("%s:%d", lib.Server.Host, lib.Server.Port)
	log.Printf(serverPort)
	r := gin.New()
	r = route.InitApiRoute(r)
	//r.Run()
	r.LoadHTMLGlob("/users/xiashuang/go/src/awesomeProject/src/htmls/**/*")
	server := endless.NewServer(serverPort, r)
	server.BeforeBegin = func(add string) {
		log.Printf("pid 是: %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}
