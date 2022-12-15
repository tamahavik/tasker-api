package database

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/tamahavik/tasker-api/pkg/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection interface {
	GetConnection() (*gorm.DB, error)
}

type ConnectionImpl struct {
}

func NewConnection() Connection {
	return &ConnectionImpl{}
}

func (c ConnectionImpl) GetConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
		viper.GetString("SSL_MODE"),
		viper.GetString("TZ"))

	db, err1 := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err1 != nil {
		return nil, err1
	}

	fmt.Println(viper.GetString("DB"))

	err2 := db.AutoMigrate(&models.User{})
	if err2 != nil {
		return nil, err2
	}

	sqlDriver, err3 := db.DB()
	if err3 != nil {
		return nil, err3
	}

	sqlDriver.SetMaxIdleConns(viper.GetInt("MAX_IDLE_CON"))
	sqlDriver.SetMaxOpenConns(viper.GetInt("MAX_OPEN_CON"))
	sqlDriver.SetConnMaxIdleTime(time.Duration(viper.GetInt("CON_MAX_IDLE_TIME")) * time.Minute)
	sqlDriver.SetConnMaxLifetime(time.Duration(viper.GetInt("CON_MAX_LIFE_TIME")) * time.Minute)

	return db, nil

}
