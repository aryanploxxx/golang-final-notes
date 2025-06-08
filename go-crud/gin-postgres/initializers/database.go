package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	// Connect to the database
	conn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	} else {
		log.Println("Connected to the database!")
	}

}
