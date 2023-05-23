package common

import (
	"admin-service/entity"
	system2 "admin-service/model/system"
	"admin-service/service/system"
	"admin-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type LoginInfo struct {
	LoginName string `json:"loginName"`
	LoginPwd  string `json:"loginPwd"`
	UserId    uint   `json:"userId"`
	NickName  string `json:"nickName"`
	UserType  uint   `json:"userType"` //1:管理员
}

func RegisterLogin(r *gin.RouterGroup) {
	r.POST("/login", Login)
	r.POST("/token_menu", TokenMenu)
	r.POST("/token_user_info", TokenUserInfo)

}

// Login 用户登录
func Login(c *gin.Context) {
	ret := new(entity.ApiRet)
	loginInfo := new(LoginInfo)
	if ret.Error = c.ShouldBindBodyWith(&loginInfo, binding.JSON); ret.Error == nil {
		switch loginInfo.UserType {
		case 1:
			user, _ := (&system.SysUserService{}).FindOne(&system2.SysUser{
				Username: loginInfo.LoginName,
				Password: loginInfo.LoginPwd,
			})
			if user != nil {
				loginInfo.UserId = user.ID
				loginInfo.NickName = user.NickName
				token, _ := TokenNext(*loginInfo)
				ret.Data = token
				ret.Ok = true
			}
		default:
		}
	}
	c.JSON(http.StatusOK, ret)
}

// TokenNext token生成
func TokenNext(loginInfo LoginInfo) (token string, err error) {
	jwt := utils.NewJWT()
	claims := jwt.CreateClaims(utils.BaseClaims{
		UserId:   loginInfo.UserId,
		NickName: loginInfo.NickName,
		UserType: loginInfo.UserType,
	})
	return jwt.CreateToken(claims)
}

// TokenMenu token菜单树权限
func TokenMenu(c *gin.Context) {
	ret := new(entity.ApiRet)
	var body struct {
		Token string `json:"token"`
	}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	jwt := utils.NewJWT()
	if customClaims, err := jwt.ParseToken(body.Token); err == nil {
		menuTree := (&system.SysBaseMenuService{}).GetMenuTreeByUserId(customClaims.UserId)
		ret.Ok = true
		ret.Data = menuTree

	} else {
		ret.Ok = false
		ret.Data = err.Error()
	}
	c.JSON(http.StatusOK, ret)
}

// TokenUserInfo token用户信息
func TokenUserInfo(c *gin.Context) {
	ret := new(entity.ApiRet)
	var body struct {
		Token string `json:"token"`
	}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	jwt := utils.NewJWT()
	if customClaims, err := jwt.ParseToken(body.Token); err == nil {
		ret.Ok = true
		ret.Data = customClaims
	} else {
		ret.Ok = false
		ret.Data = err.Error()
	}
	c.JSON(http.StatusOK, ret)
}
