package main

import (
	"log"

	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/logger"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
