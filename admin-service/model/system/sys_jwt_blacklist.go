package system

import (
	"admin-service/model/base"
)

//go:generate gcs-gen
type JwtBlacklist struct {
	base.GcsModelPrefix
	Jwt string `json:"jwt" gorm:"type:text;comment:jwt"`
	base.GcsModelSuffix
}

func (JwtBlacklist) TableName() string {
	return "sys_jwt_blacklist"
}

func (JwtBlacklist) ModelName() string {
	return "jwt-black-list"
}
