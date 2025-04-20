package main

import (
	"fmt"
	"log"

	"trainingsplan/api/lib"
	"trainingsplan/api/models"
)

func init() {
	config, loadConfigError := lib.LoadConfig(".")
	if loadConfigError != nil {
		log.Fatal("? Could not load environment variables", loadConfigError)
	}

	lib.ConnectToDatabase(config)
}

func main() {
	lib.Database.AutoMigrate(&models.Exercise{})
	fmt.Println("? Migration complete")
}
