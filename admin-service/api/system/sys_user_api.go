package system

import (
	"admin-service/entity"
	"admin-service/global"
	"admin-service/model/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api *SysUserApi) Register() {
	api.RouterGroup.POST("grantRole", api.GrantRole)
}

// GrantRole 用户进行角色授权
func (api *SysUserApi) GrantRole(c *gin.Context) {
	ret := new(entity.ApiRet)
	var reqBody struct {
		RoleId uint `json:"roleId"`
		UserId uint `json:"userId"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		return
	}
	var role system.SysRole
	var user system.SysUser
	global.GCS_DB.Take(&role, reqBody.RoleId)
	global.GCS_DB.Take(&user, reqBody.UserId)
	global.GCS_DB.Model(&user).Association("Roles").Replace(&role)
	ret.Ok = true
	ret.Data = user
	c.JSON(http.StatusOK, ret)
}
