package system

import (
	"admin-service/model/base"
)

//go:generate gcs-gen -hasPage=true
type SysBaseMenu struct {
	base.GcsModelPrefix
	ParentId  uint   `json:"parentId" gorm:"column:parent_id;comment:父菜单ID" AliasName:"父菜单ID"`
	Title     string `json:"title" gorm:"column:title;comment:菜单名" AliasName:"菜单名"`
	Icon      string `json:"icon" gorm:"column:icon;comment:菜单图标" AliasName:"菜单图标"`
	Name      string `json:"name" gorm:"column:name;comment:路由名" AliasName:"路由名"`
	Path      string `json:"path" gorm:"column:path;comment:路由path" AliasName:"路由路径"`
	Component string `json:"component" gorm:"column:component;comment:组件路径"  AliasName:"组件路径"`
	Hidden    bool   `json:"hidden" gorm:"column:hidden;comment:是否隐藏"  AliasName:"是否隐藏"`
	Sort      int    `json:"sort" gorm:"column:sort;comment:排序标记"  AliasName:"排序标记"`
	//Meta           `json:"meta" gorm:"embedded;comment:附加属性"`
	Roles    []SysRole     `json:"roles" gorm:"-"`
	Children []SysBaseMenu `json:"children" gorm:"-"`
	//BaseMenuParams []SysBaseMenuParam `json:"parameters" gorm:"foreignKey:base_menu_id"`
	//BaseMenuBtns   []SysBaseMenuBtn   `json:"menuBtn" gorm:"foreignKey:base_menu_id"`
	base.GcsModelSuffix
}

func (SysBaseMenu) TableName() string {
	return "sys_base_menu"
}

func (SysBaseMenu) ModelName() string {
	return "menu"
}
