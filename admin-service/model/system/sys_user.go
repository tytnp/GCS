package system

import (
	"admin-service/model/base"
)

//go:generate gcs-gen -hasPage=true
type SysUser struct {
	base.GcsModelPrefix
	//UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                                                     // 用户UUID
	Username      string    `json:"userName" gorm:"column:username;index;comment:用户登录名" AliasName:"登录名"`                  // 用户登录名
	Password      string    `json:"password"  gorm:"column:password;comment:用户登录密码"  AliasName:"密码"`                      // 用户登录密码
	NickName      string    `json:"nickName" gorm:"column:nick_name;default:系统用户;comment:用户昵称" AliasName:"昵称"`            // 用户昵称
	SideMode      string    `json:"sideMode" gorm:"default:dark;comment:用户侧边主题" `                                         // 用户侧边主题
	HeaderImg     string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor     string    `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor   string    `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	Phone         string    `json:"phone"  gorm:"column:phone;comment:用户手机号" AliasName:"手机"`                              // 用户手机号
	Email         string    `json:"email"  gorm:"column:email;comment:用户邮箱" AliasName:"邮箱"`                               // 用户邮箱
	HasSuperAdmin bool      `json:"hasSuperAdmin" gorm:"column:has_super_admin;default:1;comment:是否超管" AliasName:"是否超管"`
	Enable        bool      `json:"enable" gorm:"column:enable;default:1;comment:启用状态" AliasName:"启用状态"` //用户是否被冻结 1正常 2冻结
	Roles         []SysRole `json:"roles" gorm:"many2many:sys_user_role;joinForeignKey:UserId;joinReferences:RoleId"`
	base.GcsModelSuffix
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (SysUser) ModelName() string {
	return "user"
}
