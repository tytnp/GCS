package system

import (
	"admin-service/model/base"
)

//go:generate gcs-gen -hasPage=true
type SysDictionary struct {
	base.GcsModelPrefix
	Name              string                `json:"name" form:"name" gorm:"column:name;comment:字典名（中）" AliasName:"字典名(中)"` // 字典名（中）
	Type              string                `json:"type" form:"type" gorm:"column:type;comment:字典名（英）" AliasName:"字典名(英)"` // 字典名（英）
	Status            bool                  `json:"status" form:"status" gorm:"column:status;comment:状态" AliasName:"状态"`   // 状态
	Desc              string                `json:"desc" form:"desc" gorm:"column:desc;comment:描述" AliasName:"描叙"`         // 描述
	DictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" gorm:"foreignKey:dictionary_id" form:"sysDictionaryDetail"`
	base.GcsModelSuffix
}

func (SysDictionary) TableName() string {
	return "sys_dictionary"
}

func (SysDictionary) ModelName() string {
	return "dictionary"
}
