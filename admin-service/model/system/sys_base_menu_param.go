package system

import (
	"admin-service/model/base"
)

//go:generate gcs-gen
type SysBaseMenuParam struct {
	base.GcsModelPrefix
	BaseMenuID uint   `json:"baseMenuID" gorm:"comment:菜单ID"`
	Type       string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Key        string `json:"key" gorm:"comment:地址栏携带参数的key"`
	Value      string `json:"value" gorm:"comment:地址栏携带参数的值"`
	base.GcsModelSuffix
}

func (SysBaseMenuParam) TableName() string {
	return "sys_base_menu_param"
}

func (SysBaseMenuParam) ModelName() string {
	return "menu-param"
}
