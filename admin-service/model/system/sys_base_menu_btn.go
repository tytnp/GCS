package system

import (
	"admin-service/model/base"
)

//go:generate gcs-gen
type SysBaseMenuBtn struct {
	base.GcsModelPrefix
	BaseMenuID uint   `json:"baseMenuID" gorm:"comment:菜单ID"`
	Name       string `json:"name" gorm:"comment:按钮名称"`
	Desc       string `json:"desc" gorm:"comment:按钮备注"`
	base.GcsModelSuffix
}

func (SysBaseMenuBtn) TableName() string {
	return "sys_base_menu_btn"
}

func (SysBaseMenuBtn) ModelName() string {
	return "menu-btn"
}
