package system

//go:generate gcs-gen
type SysRoleMenu struct {
	RoleId uint `json:"roleId" gorm:"primaryKey;column:role_id;comment:角色ID;"`
	MenuId uint `json:"menuId" gorm:"primaryKey;column:menu_id;comment:菜单ID;"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

func (SysRoleMenu) ModelName() string {
	return "role-menu"
}
