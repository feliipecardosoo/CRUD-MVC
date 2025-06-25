package main

import (
	"crud/src/configuration/logger"
	"crud/src/controller"
	"crud/src/controller/routes"
	"crud/src/model/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Começando servidor")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Inicializa Dependencias
	service := service.NewMembroDomainService()
	userController := controller.NewMembroControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
