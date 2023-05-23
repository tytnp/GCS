package api

import (
	"admin-service/api/base"
	"admin-service/api/common"
	_ "admin-service/api/system"
	"admin-service/global"
	"admin-service/resources"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var Context = base.ApiContext

func Listening(r *gin.RouterGroup) {
	for _, value := range Context {
		value.Default(r)
		value.Register()
	}
	common.RegisterLogin(r)        //登录
	resources.EmbedAdminVuePage(r) //嵌入后台管理页面
	global.GCS_Engine.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/admin") {
			c.Data(http.StatusOK, "text/html", resources.AdminIndexFile)
		}
	})
}
