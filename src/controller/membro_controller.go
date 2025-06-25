package controller

import (
	"crud/src/model/service"

	"github.com/gin-gonic/gin"
)

func NewMembroControllerInterface(serviceInterface service.MembroDomainService) MembroControllerInterface {
	return &membroControllerInterface{
		service: serviceInterface,
	}
}

type MembroControllerInterface interface {
	AtualizaMembro(c *gin.Context)
	CreateMembroController(c *gin.Context)
	DeleteMembro(c *gin.Context)

	FindMembroById(c *gin.Context)
	FindMembroByEmail(c *gin.Context)
}

type membroControllerInterface struct {
	service service.MembroDomainService
}
