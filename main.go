package main

import (
	"crud/src/configuration/logger"
	controllerAdmin "crud/src/controller/admin-controller"
	controller "crud/src/controller/membro-controller"
	"crud/src/controller/routes"
	serviceAdmin "crud/src/model/service/admin-service"
	service "crud/src/model/service/membro-service"
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
	serviceMembro := service.NewMembroDomainService()
	serviceAdmin := serviceAdmin.NewAdminDomainService()
	membroController := controller.NewMembroControllerInterface(serviceMembro)
	adminController := controllerAdmin.NewAdminControllerInterface(serviceAdmin)

	router := gin.Default()

	routes.InitRoutesMembros(&router.RouterGroup, membroController)
	routes.InitRoutesAdmins(&router.RouterGroup, adminController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
