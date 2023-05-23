package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	MainConfig *viper.Viper
)

func LoadConfig() {
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

	switch active {
	case "dev":
		MainConfig := viper.New()
		MainConfig.AddConfigPath("./resources")
		MainConfig.SetConfigName("config.dev")
		MainConfig.SetConfigType("yaml")
	case "pro":
		MainConfig := viper.New()
		MainConfig.AddConfigPath("./resources")
		MainConfig.SetConfigName("config.pro")
		MainConfig.SetConfigType("yaml")
	default:
		fmt.Println("nonsupport")
	}
}
