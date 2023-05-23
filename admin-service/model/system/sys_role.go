package system

import (
	"admin-service/model/base"
)

//go:generate gcs-gen -hasPage=true
type SysRole struct {
	base.GcsModelPrefix
	RoleName      string        `json:"roleName" gorm:"column:role_name;comment:角色名" AliasName:"角色名"`
	DefaultRouter string        `json:"defaultRouter" gorm:"column:default_router;comment:默认菜单;default:dashboard" AliasName:"默认菜单"` // 默认菜单(默认dashboard)
	Users         []SysUser     `json:"users" gorm:"-"`
	BaseMenus     []SysBaseMenu `json:"baseMenus" gorm:"many2many:sys_role_menu;joinForeignKey:RoleId;joinReferences:MenuId"`
	base.GcsModelSuffix
}

func (m SysRole) TableName() string {
	return "sys_role"
}

func (m SysRole) ModelName() string {
	return "role"
}
