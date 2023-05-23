package base

import (
	"github.com/gin-gonic/gin"
)

var ApiContext []WebApi

type WebApi interface {
	Default(r *gin.RouterGroup)
	Register()
}
