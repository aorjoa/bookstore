package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error) {
	dsn := "host=localhost user=username password=P@ssw0rd dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
