package dao

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	conn_string := os.Getenv("DB_URL")
	D, err := gorm.Open(postgres.Open(conn_string), &gorm.Config{})
	DB = D
	return err
}