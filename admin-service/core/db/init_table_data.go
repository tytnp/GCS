package db

import (
	"admin-service/global"
	"admin-service/model/base"
	"admin-service/model/system"
)

func InitTableData() {
	user := system.SysUser{
		GcsModelPrefix: base.GcsModelPrefix{ID: 1},
		Username:       "admin",
		Password:       "admin",
		NickName:       "admin",
		HeaderImg:      "https://qmplusimg.henrongyi.top/gva_header.jpg",
		HasSuperAdmin:  true,
		Enable:         true,
	}
	global.GCS_DB.FirstOrCreate(&user)

	role := system.SysRole{
		GcsModelPrefix: base.GcsModelPrefix{ID: 1},
		RoleName:       "管理员",
		DefaultRouter:  "/system/role",
	}
	global.GCS_DB.FirstOrCreate(&role)

	UserRole := system.SysUserRole{
		UserId: 1,
		RoleId: 1,
	}
	global.GCS_DB.FirstOrCreate(&UserRole)

	Menu1 := system.SysBaseMenu{
		GcsModelPrefix: base.GcsModelPrefix{ID: 1},
		ParentId:       0,
		Title:          "系统管理",
		Icon:           "SettingsSharp",
		Name:           "system",
		Path:           "/system",
		Component:      "views/index.vue",
		Hidden:         false,
	}
	global.GCS_DB.FirstOrCreate(&Menu1)

	Menu2 := system.SysBaseMenu{
		GcsModelPrefix: base.GcsModelPrefix{ID: 2},
		ParentId:       1,
		Title:          "用户管理",
		Icon:           "PeopleFilled",
		Name:           "user",
		Path:           "/system/user",
		Component:      "views/system/user/index.vue",
		Hidden:         false,
	}
	global.GCS_DB.FirstOrCreate(&Menu2)

	Menu3 := system.SysBaseMenu{
		GcsModelPrefix: base.GcsModelPrefix{ID: 3},
		ParentId:       1,
		Title:          "角色管理",
		Icon:           "VerifiedUserRound",
		Name:           "role",
		Path:           "/system/role",
		Component:      "views/system/role/index.vue",
		Hidden:         false,
	}
	global.GCS_DB.FirstOrCreate(&Menu3)

	Menu4 := system.SysBaseMenu{
		GcsModelPrefix: base.GcsModelPrefix{ID: 4},
		ParentId:       1,
		Title:          "菜单管理",
		Icon:           "MenuBookSharp",
		Name:           "menu",
		Path:           "/system/menu",
		Component:      "views/system/menu/index.vue",
		Hidden:         false,
	}
	global.GCS_DB.FirstOrCreate(&Menu4)

	Menu5 := system.SysBaseMenu{
		GcsModelPrefix: base.GcsModelPrefix{ID: 5},
		ParentId:       1,
		Title:          "接口管理",
		Icon:           "ConnectingAirportsOutlined",
		Name:           "api",
		Path:           "/system/api",
		Component:      "views/system/api/index.vue",
		Hidden:         false,
	}
	global.GCS_DB.FirstOrCreate(&Menu5)

	Menu6 := system.SysBaseMenu{
		GcsModelPrefix: base.GcsModelPrefix{ID: 6},
		ParentId:       1,
		Title:          "字典管理",
		Icon:           "DragIndicatorSharp",
		Name:           "dictionary",
		Path:           "/system/dictionary",
		Component:      "views/system/dictionary/index.vue",
		Hidden:         false,
	}
	global.GCS_DB.FirstOrCreate(&Menu6)

	Menu7 := system.SysBaseMenu{
		GcsModelPrefix: base.GcsModelPrefix{ID: 7},
		ParentId:       1,
		Title:          "字典详情管理",
		Icon:           "DragIndicatorSharp",
		Name:           "dictionary-detail",
		Path:           "/system/dictionary-detail/:id",
		Component:      "views/system/dictionary-detail/index.vue",
		Hidden:         true,
	}
	global.GCS_DB.FirstOrCreate(&Menu7)

	RoleMenu1 := system.SysRoleMenu{
		RoleId: 1,
		MenuId: 1,
	}
	global.GCS_DB.FirstOrCreate(&RoleMenu1)

	RoleMenu2 := system.SysRoleMenu{
		RoleId: 1,
		MenuId: 2,
	}
	global.GCS_DB.FirstOrCreate(&RoleMenu2)

	RoleMenu3 := system.SysRoleMenu{
		RoleId: 1,
		MenuId: 3,
	}
	global.GCS_DB.FirstOrCreate(&RoleMenu3)

	RoleMenu4 := system.SysRoleMenu{
		RoleId: 1,
		MenuId: 4,
	}
	global.GCS_DB.FirstOrCreate(&RoleMenu4)

	RoleMenu5 := system.SysRoleMenu{
		RoleId: 1,
		MenuId: 5,
	}
	global.GCS_DB.FirstOrCreate(&RoleMenu5)

	RoleMenu6 := system.SysRoleMenu{
		RoleId: 1,
		MenuId: 6,
	}
	global.GCS_DB.FirstOrCreate(&RoleMenu6)

	RoleMenu7 := system.SysRoleMenu{
		RoleId: 1,
		MenuId: 7,
	}
	global.GCS_DB.FirstOrCreate(&RoleMenu7)
}
