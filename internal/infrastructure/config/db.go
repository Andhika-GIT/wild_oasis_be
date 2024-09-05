package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper, logger *logger.Interface) *gorm.DB{
	DB_USERNAME := viper.GetString("DB_USERNAME")
	DB_PASSWORD:= viper.GetString("DB_PASSWORD")
	DB_HOST := viper.GetString("DB_HOST")
	DB_PORT := viper.GetString("DB_PORT")
	DB_NAME := viper.GetString("DB_NAME")
	DB_IDLE_CONNECTION := viper.GetInt("DB_IDLE_CONNECTION")
	DB_MAX_CONNECTION := viper.GetInt("DB_MAX_CONNECTION")
	DB_MAX_LIFE_TIME := viper.GetInt("DB_MAX_LIFE_TIME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: *logger,
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(DB_IDLE_CONNECTION)
	connection.SetMaxOpenConns(DB_MAX_CONNECTION)
	connection.SetConnMaxLifetime(time.Second * time.Duration(DB_MAX_LIFE_TIME))

	return db
}