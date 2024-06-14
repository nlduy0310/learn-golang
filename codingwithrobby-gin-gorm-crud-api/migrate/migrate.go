package main

import (
	"example/gin-gorm-crud/initializers"
	"example/gin-gorm-crud/models"
	"log"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("An error happened when migrating: ", err)
	} else {
		log.Print("Migration completed successfully")
	}
}
