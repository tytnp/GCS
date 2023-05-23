package system

import api "admin-service/api/base"
import model "admin-service/model/system"
import service "admin-service/service/system"

type SysUserApi struct {
	api.Api[model.SysUser]
	service.SysUserService
}

func init() {
	api.ApiContext = append(api.ApiContext, &SysUserApi{})
	api.ApiContext = append(api.ApiContext, &SysApiApi{})
	api.ApiContext = append(api.ApiContext, &SysBaseMenuApi{})
	api.ApiContext = append(api.ApiContext, &SysBaseMenuBtnApi{})
	api.ApiContext = append(api.ApiContext, &SysBaseMenuParamApi{})
	api.ApiContext = append(api.ApiContext, &SysDictionaryApi{})
	api.ApiContext = append(api.ApiContext, &SysDictionaryDetailApi{})
	api.ApiContext = append(api.ApiContext, &JwtBlacklistApi{})
	api.ApiContext = append(api.ApiContext, &SysRoleApi{})
	api.ApiContext = append(api.ApiContext, &SysRoleMenuApi{})
	api.ApiContext = append(api.ApiContext, &SysUserRoleApi{})
}

type SysApiApi struct {
	api.Api[model.SysApi]
	service.SysApiService
}
type SysBaseMenuApi struct {
	api.Api[model.SysBaseMenu]
	service.SysBaseMenuService
}
type SysBaseMenuBtnApi struct {
	api.Api[model.SysBaseMenuBtn]
	service.SysBaseMenuBtnService
}
type SysBaseMenuParamApi struct {
	api.Api[model.SysBaseMenuParam]
	service.SysBaseMenuParamService
}
type SysDictionaryApi struct {
	api.Api[model.SysDictionary]
	service.SysDictionaryService
}
type SysDictionaryDetailApi struct {
	api.Api[model.SysDictionaryDetail]
	service.SysDictionaryDetailService
}
type JwtBlacklistApi struct {
	api.Api[model.JwtBlacklist]
	service.JwtBlacklistService
}
type SysRoleApi struct {
	api.Api[model.SysRole]
	service.SysRoleService
}
type SysRoleMenuApi struct {
	api.Api[model.SysRoleMenu]
	service.SysRoleMenuService
}
type SysUserRoleApi struct {
	api.Api[model.SysUserRole]
	service.SysUserRoleService
}
