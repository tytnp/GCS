package config

type SysConf struct {
	DbType       string  `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	RouterPrefix string  `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	JwtConf      JwtConf `mapstructure:"jwt_conf" json:"jwt_conf" yaml:"jwt_conf"`
}

type JwtConf struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`    // jwt签名
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
}
