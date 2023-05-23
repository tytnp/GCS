package config

import (
	"admin-service/global"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Viper() {
	config := viper.New()
	config.AddConfigPath("./resources")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	err := config.ReadInConfig()
	if err != nil {
		fmt.Println("'config.yaml' file read error:", err)
		os.Exit(0)
	}
	active := config.GetString("profiles.active")

	activeConfig := viper.New()
	activeConfig.AddConfigPath("./resources")
	activeConfig.SetConfigType("yaml")
	switch active {
	case "dev":
		activeConfig.SetConfigName("config.dev")
	case "pro":
		activeConfig.SetConfigName("config.pro")
	default:
		fmt.Println("nonsupport")
	}
	err = activeConfig.ReadInConfig()
	if err != nil {
		fmt.Println("'config."+active+".yaml' file read error:", err)
		os.Exit(0)
	}

	activeConfig.SetDefault("sys-conf.router-prefix", "")
	activeConfig.SetDefault("sys-conf.db-type", "mysql")

	activeConfig.SetDefault("mysql.engine", "InnoDB")
	//activeConfig.SetDefault("mysql.table_prefix", "gcs")
	activeConfig.SetDefault("mysql.singular_table", "true")
	activeConfig.SetDefault("mysql.log-mode", "Info")
	if err = activeConfig.Unmarshal(&global.GCS_Conf); err != nil {
		fmt.Println(err)
	}
	global.GCS_Viper = activeConfig
}
