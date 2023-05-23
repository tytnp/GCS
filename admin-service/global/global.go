package global

import (
	"admin-service/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	GCS_Viper               *viper.Viper
	GCS_Conf                config.ActiveConf
	GCS_DB                  *gorm.DB
	GCS_Engine              *gin.Engine
	GCS_Concurrency_Control = &singleflight.Group{}
)
