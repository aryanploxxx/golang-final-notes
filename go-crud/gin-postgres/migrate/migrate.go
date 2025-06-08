package main

import (
	"gin-postgres/initializers"
	"gin-postgres/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	// AutoMigrate automatically creates the table for the Post model if it doesn't exist.
	// It also creates the table if the model has changed.
	// go run migrate/migrate.go
}
