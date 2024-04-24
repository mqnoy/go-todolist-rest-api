package config

import (
	"github.com/mqnoy/go-todolist-rest-api/pkg/clogger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMySQLDatabase(cfg Configuration) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.Database.Mysql.DSN()), &gorm.Config{
		Logger: gormlogger.New(
			clogger.Logger().WithField("location", "database"),
			gormlogger.Config{
				Colorful:             false,
				LogLevel:             gormlogger.LogLevel(AppConfig.Database.LogLevel),
				ParameterizedQueries: false,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
	})

	if err != nil {
		clogger.Logger().SetReportCaller(true)
		clogger.Logger().Fatalf("Error creating database connection err: %v", err)
	}

	clogger.Logger().Debugf("Database is connected to %s:%d", cfg.Database.Mysql.Host, cfg.Database.Mysql.Port)

	return db
}
