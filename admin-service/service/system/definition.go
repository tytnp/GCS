package system

import service "admin-service/service/base"
import model "admin-service/model/system"

type SysApiService struct {
	service.BaseService[model.SysApi]
}
type SysBaseMenuService struct {
	service.BaseService[model.SysBaseMenu]
}
type SysBaseMenuBtnService struct {
	service.BaseService[model.SysBaseMenuBtn]
}
type SysBaseMenuParamService struct {
	service.BaseService[model.SysBaseMenuParam]
}
type SysDictionaryService struct {
	service.BaseService[model.SysDictionary]
}
type SysDictionaryDetailService struct {
	service.BaseService[model.SysDictionaryDetail]
}
type JwtBlacklistService struct {
	service.BaseService[model.JwtBlacklist]
}
type SysRoleService struct {
	service.BaseService[model.SysRole]
}
type SysRoleMenuService struct {
	service.BaseService[model.SysRoleMenu]
}
type SysUserService struct {
	service.BaseService[model.SysUser]
}
type SysUserRoleService struct {
	service.BaseService[model.SysUserRole]
}
