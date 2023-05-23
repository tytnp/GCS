package main

import (
	"admin-service/core/config"
	"admin-service/core/db"
	"admin-service/core/web"
	"admin-service/global"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	config.Viper()
	db.Gorm()
	if global.GCS_DB != nil {
		db, _ := global.GCS_DB.DB()
		defer db.Close()
	}
	web.StartGinServer()
}
