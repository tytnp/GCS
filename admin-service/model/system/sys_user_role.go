package system

//go:generate gcs-gen
type SysUserRole struct {
	UserId uint `json:"userId" gorm:"primaryKey;column:user_id;comment:用户ID;"`
	RoleId uint `json:"roleId" gorm:"primaryKey;column:role_id;comment:角色ID;"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}

func (SysUserRole) ModelName() string {
	return "user-role"
}
