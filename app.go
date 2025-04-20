package main

import (
	"log"
	"trainingsplan/api/lib"
	"trainingsplan/api/routers/exercises"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "trainingsplan/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Swagger Example API
// @Version 1.0
// @Description Another Trainingsplan App
func main() {
	config, loadConfigError := lib.LoadConfig(".")
	if loadConfigError != nil {
		log.Fatal("Could not load environment variables")
	}

	lib.ConnectToDatabase(config)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.ClientOrigin}
	router.Use(cors.New(corsConfig))

	exercises.CreateRouter(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(router.Run(":" + config.ServerPort))
}
