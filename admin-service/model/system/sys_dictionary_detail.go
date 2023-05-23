package system

import (
	"admin-service/model/base"
)

//go:generate gcs-gen -hasPage=true
type SysDictionaryDetail struct {
	base.GcsModelPrefix
	DictionaryID int    `json:"dictionaryID" form:"SysDictionary" gorm:"column:dictionary_id;comment:关联标记"` // 关联标记
	Label        string `json:"label" form:"label" gorm:"column:label;comment:展示值" AliasName:"展示值"`         // 展示值
	Value        int    `json:"value" form:"value" gorm:"column:value;comment:字典值" AliasName:"字典值"`         // 字典值
	Status       bool   `json:"status" form:"status" gorm:"column:status;comment:启用状态" AliasName:"状态"`      // 启用状态
	Sort         int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记" AliasName:"排序"`            // 排序标记
	base.GcsModelSuffix
}

func (SysDictionaryDetail) TableName() string {
	return "sys_dictionary_detail"
}

func (SysDictionaryDetail) ModelName() string {
	return "dictionary-detail"
}
