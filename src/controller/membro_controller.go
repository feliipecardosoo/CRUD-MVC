package controller

import "github.com/gin-gonic/gin"

type MembroCOntrollerInterface interface {
	AtualizaMembro(c *gin.Context)
	CreateMembroController(c *gin.Context)
	DeleteMembro(c *gin.Context)

	FindMembroById(c *gin.Context)
	FindMembroByEmail(c *gin.Context)
}
