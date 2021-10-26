package lib

import (
	"fmt"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/config"
)

var (
	Minprogram *miniprogram.MiniProgram
)

func init() {
	initMiniprogram()
}

func initMiniprogram() {
	obj := wechat.NewWechat()
	newRedis := cache.NewRedis(&cache.RedisOpts{
		Host:        fmt.Sprintf("%s:%s", RedisConf.Host, RedisConf.Port),
		Password:    RedisConf.password,
		Database:    0,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60,
	})
	cfg := &config.Config{
		AppID:     MpConf.AppId,
		AppSecret: MpConf.AppSecret,
		Cache:     newRedis,
	}
	Minprogram = obj.GetMiniProgram(cfg)
}
