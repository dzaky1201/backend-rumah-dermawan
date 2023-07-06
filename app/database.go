package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=rumah_dermawan port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
