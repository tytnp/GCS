package web

import (
	"admin-service/api"
	"admin-service/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartGinServer() {
	//gin.SetMode(gin.ReleaseMode) //gin模式
	Engine := gin.Default()
	global.GCS_Engine = Engine
	Engine.Use(cors.Default()) //允许跨域
	RouterGroup := Engine.Group(global.GCS_Conf.SysConf.RouterPrefix)
	api.Listening(RouterGroup)
	Engine.Run(":6677")
}
