package config

import (
	"admin-service/global"
	"admin-service/resources"
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Viper() {
	config := viper.New()
	configFile, err := resources.ConfigFs.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("'config.yaml' file read error:", err)
		os.Exit(0)
	}
	config.SetConfigType("yaml")
	err = config.ReadConfig(bytes.NewBuffer(configFile))
	if err != nil {
		fmt.Println("'config.yaml' file parser error:", err)
		os.Exit(0)
	}
	active := config.GetString("profiles.active")

	activeConfig := viper.New()
	activeConfig.SetConfigType("yaml")
	switch active {
	case "dev":
		activeConfigFile, _ := resources.ConfigFs.ReadFile("config.dev.yaml")
		err = activeConfig.ReadConfig(bytes.NewBuffer(activeConfigFile))
	case "pro":
		activeConfigFile, _ := resources.ConfigFs.ReadFile("config.pro.yaml")
		err = activeConfig.ReadConfig(bytes.NewBuffer(activeConfigFile))
	default:
		fmt.Println("nonsupport")
	}
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
