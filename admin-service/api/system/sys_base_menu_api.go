package system

import (
	"admin-service/entity"
	"admin-service/model/system"
	system2 "admin-service/service/system"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func (api *SysBaseMenuApi) Register() {
	api.RouterGroup.POST("GetMenuTree", api.GetMenuTree)
	api.RouterGroup.POST("GetMenuTreeByRole", api.GetMenuTreeByRole)
}

// GetMenuTree 获得整个菜单树
func (api *SysBaseMenuApi) GetMenuTree(c *gin.Context) {
	model := new(system.SysBaseMenu)
	treeRoot := &system.SysBaseMenu{}
	//global.GCS_DB.Find(data, model)
	if data, _, err := api.FindList(model); err == nil {
		treeRoot.ID = 0
		treeRoot.Title = "菜单根"
		treeRoot.Path = "/"
		treeRoot.Name = "root"
		//fmt.Println("总数量:", totalCount)
		api.SysBaseMenuService.FormatMenuTree(treeRoot, data)
		//api.Service.FormatMenuTree()
	}
	c.JSON(http.StatusOK, treeRoot)
}

// GetMenuTreeByRole 获得获得角色下菜单树
func (api *SysBaseMenuApi) GetMenuTreeByRole(c *gin.Context) {
	ret := new(entity.ApiRet)
	var role system.SysRole
	if ret.Error = c.ShouldBindBodyWith(&role, binding.JSON); ret.Error == nil {
		ret.Ok = true
		ret.Data = (&system2.SysBaseMenuService{}).GetMenuByRoleId(role.ID)
	}
	c.JSON(http.StatusOK, ret)
}

func (api *SysBaseMenuApi) FormatMenuTree(menu *system.SysBaseMenu, data *[]system.SysBaseMenu) {
	for i := 0; i < len(*data); i++ {
		if menu.ID == (*data)[i].ParentId {
			//fmt.Println(menu.Title, (*data)[i].Title)
			menu.Children = append(menu.Children, (*data)[i])
			(*data) = append((*data)[:i], (*data)[i+1:]...)
			i--
		}
	}
	for i := 0; i < len(menu.Children); i++ {
		api.FormatMenuTree(&menu.Children[i], data)
	}
}
