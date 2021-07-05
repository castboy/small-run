package libraries

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"
	"time"
)

func NewDB(v *viper.Viper, lc fx.Lifecycle, log *log.Logger) (db *gorm.DB, err error) {
	config := v.Sub("mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetString("user"),
		config.GetString("password"),
		config.GetString("host"),
		config.GetInt("port"),
		config.GetString("dbname"),
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	sqlDB.SetConnMaxLifetime(config.GetDuration("connMaxLifetime") * time.Minute)
	sqlDB.SetMaxIdleConns(config.GetInt("maxIdleConns"))
	sqlDB.SetMaxOpenConns(config.GetInt("maxOpenConns"))

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("close sql conn.")
			return sqlDB.Close()
		},
	})

	return
}
