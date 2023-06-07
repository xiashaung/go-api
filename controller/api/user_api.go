package api

import (
	"github.com/gin-gonic/gin"
	"go-api/lib"
	"go-api/model"
	"go-api/service"
)

func init() {
	c := &UserApi{}
	c.init(lib.R) //这里需要引入你的gin框架的实例
}

func (t UserApi) init(g *gin.Engine) {
	// 依次: 分页列表，单条，插入，修改，删除
	group := g.Group("/user")
	group.GET("/list", t.list) //不设置限制条件的画默认查询所有
	group.GET("/one", t.one)
	group.POST("/insert", t.insert)
	group.POST("/update", t.update)
	group.POST("/delete", t.delete)
}

// UserApi 控制器
type UserApi struct {
	Page int   `form:"page,default=1"`
	Size int   `form:"size,default=10"`
	Ids  []int `form:"ids"`
}

// 分页列表
func (t UserApi) list(c *gin.Context) {
	_ = c.Bind(&t)
	v := model.User{}
	_ = c.Bind(&v)
	lib.Resp.SuccessJson(c, service.UserService.List(t.Page, t.Size, &v))
}

// 根据主键Id查询记录
func (t UserApi) one(c *gin.Context) {
	var v model.User
	_ = c.Bind(&v)
	lib.Resp.SuccessJson(c, service.UserService.One(v.Id))
}

// 修改记录
func (t UserApi) update(c *gin.Context) {
	var v model.User
	_ = c.Bind(&v)
	service.UserService.Update(v)
	lib.Resp.SuccessMsg(c, "修改成功！")
}

// 插入记录
func (t UserApi) insert(c *gin.Context) {
	var v model.User
	_ = c.Bind(&v)
	user, _ := service.UserService.FindByMobile(v.Mobile)
	if user.Id > 0 {
		lib.Resp.ErrorJson(c, 101, "该用户已存在")
		return
	}
	service.UserService.Insert(v)
	lib.Resp.SuccessMsg(c, "插入成功！")
}

// 根据主键删除
func (t UserApi) delete(c *gin.Context) {
	_ = c.Bind(&t)
	service.UserService.Delete(t.Ids)
	lib.Resp.SuccessMsg(c, "删除成功！")
}
