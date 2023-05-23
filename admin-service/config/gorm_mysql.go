package config

type Mysql struct {
	Host          string `mapstructure:"host" json:"host" yaml:"host"`
	Port          string `mapstructure:"port" json:"port" yaml:"port"`
	Username      string `mapstructure:"username" json:"username" yaml:"username"`
	Password      string `mapstructure:"password" json:"password" yaml:"password"`
	DbName        string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Param         string `mapstructure:"param" json:"param" yaml:"param"`                                           //数据库高级参数配置
	Engine        string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`                       //数据库引擎，默认InnoDB
	TablePrefix   string `mapstructure:"table_prefix" json:"table_prefix" yaml:"table_prefix" default:"gcs"`        //全局表前缀，单独定义TableName则不生效,默认gcs
	SingularTable bool   `mapstructure:"singular_table" json:"singular_table" yaml:"singular_table" default:"true"` //是否开启全局禁用复数，true表示开启
	LogMode       string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                                  //日志模式 Silent Error Error Info
	MaxIdleConns  int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`                //空闲中的最大连接数
	MaxOpenConns  int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`                //打开的最大连接数
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DbName + "?" + m.Param
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
