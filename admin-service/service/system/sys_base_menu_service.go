package system

import (
	"admin-service/global"
	"admin-service/model/system"
)

// GetMenuTreeByUserId 根据用户权限获得菜单树
func (s *SysBaseMenuService) GetMenuTreeByUserId(userId uint) (menuTree *system.SysBaseMenu) {
	user := new(system.SysUser)
	menuTree = &system.SysBaseMenu{}
	global.GCS_DB.Preload("Roles").Take(&user, userId)
	if len(user.Roles) > 0 {
		role := user.Roles[0]
		menuTree.ID = 0
		menuTree.Title = "菜单根"
		menuTree.Path = role.DefaultRouter
		menuTree.Name = "root"
		if user.HasSuperAdmin {
			if data, _, err := s.FindList(&system.SysBaseMenu{}); err == nil {
				s.FormatMenuTree(menuTree, data)
			}
		} else {
			menus := s.GetMenuByRoleId(role.ID)
			s.FormatMenuTree(menuTree, &menus)
		}
	}
	return menuTree
}

// GetMenuByRoleId 根据角色权限获得菜单
func (s *SysBaseMenuService) GetMenuByRoleId(roleId uint) []system.SysBaseMenu {
	var role system.SysRole
	global.GCS_DB.Preload("BaseMenus").Take(&role, roleId)
	return role.BaseMenus
}

// FormatMenuTree 格式化菜单树
func (s *SysBaseMenuService) FormatMenuTree(menu *system.SysBaseMenu, data *[]system.SysBaseMenu) {
	for i := 0; i < len(*data); i++ {
		if menu.ID == (*data)[i].ParentId {
			//fmt.Println(menu.Title, (*data)[i].Title)
			menu.Children = append(menu.Children, (*data)[i])
			(*data) = append((*data)[:i], (*data)[i+1:]...)
			i--
		}
	}
	for i := 0; i < len(menu.Children); i++ {
		s.FormatMenuTree(&menu.Children[i], data)
	}
}
