package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/fourtf/studyhub/models"
	"github.com/jinzhu/gorm"

	//Needed to use the postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//ConnectToDB establishes a connection to the database and handles errors if the database is not available
func ConnectToDB() (*gorm.DB, error) {

	username := os.Getenv("databaseUser")
	password := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	databaseType := os.Getenv("databaseType")

	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", username, databaseName, password)

	db, err := gorm.Open(databaseType, connectionString)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	log.Println("Successfully connected to the database!")
	return db, nil
}
