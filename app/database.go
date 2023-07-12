package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDb() *gorm.DB {
	// load env variables
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")
	// dbPort := os.Getenv("DB_PORT")
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

	dsn := "host=127.0.0.1 user=rumahder_admin password=rumahdermawan12 dbname=rumahder_backend port=5432 sslmode=enable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
