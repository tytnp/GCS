package db

import (
	"admin-service/global"
	"admin-service/model/system"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Gorm() {
	switch global.GCS_Conf.SysConf.DbType {
	case "mysql":
		gormMysql()
	case "sqlite":
		gormSqlite()
	}
	if global.GCS_DB != nil {
		InitializeTables()
	}
}

func InitializeTables() {
	db := global.GCS_DB
	db.SetupJoinTable(&system.SysRole{}, "BaseMenus", &system.SysRoleMenu{})
	db.SetupJoinTable(&system.SysUser{}, "Roles", &system.SysUserRole{})
	db.AutoMigrate(
		system.SysUser{},
		system.SysRole{},
		system.SysBaseMenu{},
		system.SysApi{},

		system.SysDictionary{},
		system.SysDictionaryDetail{},

		//system.SysBaseMenuBtn{},
		//system.SysBaseMenuParam{},

		//system.SysApi{},
		//

		//system.JwtBlacklist{},
	)
	InitTableData()
}

func gormSqlite() {
	db, err := gorm.Open(sqlite.Open("gcs.db"), &gorm.Config{
		//Logger:      logger.Default.LogMode(logger.Info),
		QueryFields: true,
	})
	if err != nil {
		panic("failed to connect database")
	} else {
		global.GCS_DB = db
	}
}

func gormMysql() {
	mysqlConfg := global.GCS_Conf.Mysql
	dialector := mysql.Config{
		DSN:                       mysqlConfg.Dsn(),
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
	}
	//数据库配置
	opts := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mysqlConfg.TablePrefix,
			SingularTable: mysqlConfg.SingularTable,
		},
	}
	//设置日志级别
	switch global.GCS_Conf.Mysql.LogMode {
	case "silent", "Silent":
		opts.Logger = logger.Default.LogMode(logger.Silent)
	case "error", "Error":
		opts.Logger = logger.Default.LogMode(logger.Error)
	case "warn", "Warn":
		opts.Logger = logger.Default.LogMode(logger.Warn)
	case "info", "Info":
		opts.Logger = logger.Default.LogMode(logger.Info)
	default:
		opts.Logger = logger.Default.LogMode(logger.Info)
	}
	//opts.Logger = logger
	if db, err := gorm.Open(mysql.New(dialector), opts); err != nil {
		return
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+mysqlConfg.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(mysqlConfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(mysqlConfg.MaxOpenConns)
		global.GCS_DB = db
	}
}
