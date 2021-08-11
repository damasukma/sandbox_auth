package utils

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Connection struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
)

func NewConnectionDB(connection *Connection) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", connection.User, connection.Password, connection.Host, connection.Port, connection.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	return db, nil
}
