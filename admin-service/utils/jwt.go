package utils

import (
	"admin-service/global"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JWT struct {
	SigningKey []byte
}

type BaseClaims struct {
	UserId   uint   `json:"userId"`
	NickName string `json:"nickName"`
	UserType uint   `json:"userType"` //1:管理员
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GCS_Conf.SysConf.JwtConf.SigningKey),
	}
}

// 创建一个Claims
func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	bufferTime, _ := ParseDuration(global.GCS_Conf.SysConf.JwtConf.BufferTime)
	expiresTime, _ := ParseDuration(global.GCS_Conf.SysConf.JwtConf.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bufferTime / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                         // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),       // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresTime)), // 过期时间 7天  配置文件
			Issuer:    global.GCS_Conf.SysConf.JwtConf.Issuer,          // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string) (string, error) {
	claims, err := j.ParseToken(oldToken)
	if err == nil {
		expiresTime, _ := ParseDuration(global.GCS_Conf.SysConf.JwtConf.ExpiresTime)
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expiresTime))
		v, err, _ := global.GCS_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
			return j.CreateToken(*claims)
		})
		return v.(string), err
	}
	return "", err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, nil
	} else {
		return nil, nil
	}
}
