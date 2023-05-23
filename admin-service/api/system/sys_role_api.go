package system

import (
	"admin-service/entity"
	"admin-service/global"
	"admin-service/model/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api *SysRoleApi) Register() {
	api.RouterGroup.POST("grantMenu", api.GrantMenu)
}

// GrantMenu 角色进行菜单授权
func (api *SysRoleApi) GrantMenu(c *gin.Context) {
	ret := new(entity.ApiRet)
	var reqBody struct {
		RoleId  uint   `json:"roleId"`
		MenuIds []uint `json:"menuIds"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		return
	}
	var role system.SysRole
	var menus []system.SysBaseMenu
	global.GCS_DB.Take(&role, reqBody.RoleId)
	global.GCS_DB.Find(&menus, reqBody.MenuIds)
	global.GCS_DB.Model(&role).Association("BaseMenus").Replace(&menus)
	ret.Ok = true
	ret.Data = role
	c.JSON(http.StatusOK, ret)
}
