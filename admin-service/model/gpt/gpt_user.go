package gpt

import (
	"admin-service/model/base"
)

//go:generate gcs-gen -hasPage=true
type GptUser struct {
	base.GcsModelPrefix
	DevId      string `json:"devId" gorm:"column:devId;index;comment:设备编号" AliasName:"设备编号"`              // 设备编号
	Username   string `json:"userName" gorm:"column:username;index;comment:用户登录名" AliasName:"登录名"`        // 用户登录名
	Password   string `json:"password"  gorm:"column:password;comment:用户登录密码"  AliasName:"密码"`            // 用户登录密码
	NickName   string `json:"nickName" gorm:"column:nick_name;comment:用户昵称" AliasName:"昵称"`               // 用户昵称
	InviteCode string `json:"inviteCode" gorm:"column:invite_code;default:1;comment:邀请码" AliasName:"邀请码"` //邀请码
	Enable     bool   `json:"enable" gorm:"column:enable;default:1;comment:启用状态" AliasName:"启用状态"`        //用户是否被冻结 1正常 2冻结
	base.GcsModelSuffix
}

func (GptUser) TableName() string {
	return "gpt_user"
}

func (GptUser) ModelName() string {
	return "gpt-user"
}
