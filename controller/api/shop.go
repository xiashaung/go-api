package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"go-api/lib"
	"go-api/service"
)

func GetShopInfo(c *gin.Context) {
	shopId, _ := strconv.Atoi(c.Query("shop_id"))
	shopInfoService := service.ShopInfo{}
	shopInfo := shopInfoService.GetById(shopId)
	if shopInfo.ShopID > 0 {
		lib.SuccessJson(c, shopInfo)
	} else {
		lib.SuccessJson(c, gin.H{})
	}
}
