package utils

import (
	"log"

	"github.com/fourtf/studyhub/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//ConnectToDB establishes a connection to the database and handles errors if the database is not available
func ConnectToDB() *gorm.DB {

	connectionString := "user=postgres dbname=postgres password=studyhub_dev sslmode=disable"

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		log.Println("error", err)
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.User{})

	log.Println("Successfully connected to the database!")
	return db
}
