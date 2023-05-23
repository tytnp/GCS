package config

type ActiveConf struct {
	SysConf SysConf `mapstructure:"sys-conf" json:"sys-conf" yaml:"sys-conf"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
