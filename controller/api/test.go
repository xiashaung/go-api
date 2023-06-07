package api

import (
	"github.com/gin-gonic/gin"
	"go-api/lib"
	"go-api/model"
	"go-api/service"
	"log"
)

func init() {
	c := &TestApi{}
	c.init(lib.R) //这里需要引入你的gin框架的实例
}

func (t TestApi) init(g *gin.Engine) {
	// 依次: 分页列表，单条，插入，修改，删除
	group := g.Group("/test")
	group.GET("/t", t.test)   //不设置限制条件的画默认查询所有
	group.GET("/d", t.decode) //不设置限制条件的画默认查询所有
}

// UserApi 控制器
type TestApi struct {
	Page int   `form:"page,default=1"`
	Size int   `form:"size,default=10"`
	Ids  []int `form:"ids"`
}

func (t TestApi) test(c *gin.Context) {
	var user model.User
	user.Id = 1
	service.UserService.One(user.Id)
	var expire = 7200
	token := lib.Jwt.Encode(int(user.Id), int64(expire))
	lib.Resp.SuccessJson(c, token)
}

func (t TestApi) decode(c *gin.Context) {
	token, _ := c.GetQuery("token")
	log.Println("token", token)
	lib.Resp.SuccessJson(c, lib.Jwt.Decode(token))
}
