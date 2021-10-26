package route

import (
	"github.com/gin-gonic/gin"
	"go-api/controller/api"
	"net/http"
)

func InitApiRoute(r *gin.Engine) {
	//时间api
	timeApi := r.Group("/time")
	timeApi.POST("/timestamp", api.ToTimestamp)
	timeApi.GET("/now", api.Nowtime)
	timeApi.POST("/datetime", api.ToDatetime)

	//微信授权接口
	mpApi := r.Group("/account/")
	mpApi.POST("wx-auth", api.MpAuth)

	r.GET("/queue/producer", api.QueueProducer)
	r.GET("/shop/info", api.GetShopInfo)
	r.GET("/queue/test", api.QueueTest)
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK,"这是首页")
	})
}
