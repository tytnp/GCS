package system

import (
	"admin-service/model/base"
)

//go:generate gcs-gen -hasPage=true
type SysApi struct {
	base.GcsModelPrefix
	Path        string `json:"path" gorm:"column:path;comment:api路径" AliasName:"Api路径"`                 // api路径
	Description string `json:"description" gorm:"column:description;comment:api中文描述" AliasName:"api描述"` // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"column:api_group;comment:api组" AliasName:"api组"`          // api组
	Method      string `json:"method" gorm:"column:method;default:POST;comment:方法" AliasName:"api方法"`   // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
	base.GcsModelSuffix
}

func (SysApi) TableName() string {
	return "sys_api"
}

func (SysApi) ModelName() string {
	return "api"
}
