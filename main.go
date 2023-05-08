package main

import (
	"flag"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go-api/lib"
	"go-api/route"
	"log"
	"path"
	"runtime"
	"syscall"
)

func setParams() {
	appPath := getAppPath()
	var config_path = flag.String("config_path", appPath+"/etc/config.ini", "指定配置文件目录")
	var app_path = flag.String("app_path", appPath, "指定app目录")
	flag.Parse()
	lib.ConfigPath = *config_path
	lib.APP_APTH = *app_path
}

func main() {
	setParams()
	gin.SetMode(lib.Server.Model)
	serverPort := fmt.Sprintf("%s:%d", lib.Server.Host, lib.Server.Port)
	log.Printf(serverPort)
	route.InitApiRoute(lib.R)
	//r.Run()
	//r.LoadHTMLGlob("/users/xiashuang/go/src/awesomeProject/src/htmls/**/*")
	server := endless.NewServer(serverPort, lib.R)
	server.BeforeBegin = func(add string) {
		log.Printf("pid 是: %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}

func getAppPath() string {
	var appPath string
	_, filname, _, ok := runtime.Caller(0)
	if ok {
		appPath = path.Dir(filname)
	}
	return appPath
}
