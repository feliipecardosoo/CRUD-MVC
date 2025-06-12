package controller

import (
	"crud/src/configuration/logger"
	"crud/src/configuration/validation"
	"crud/src/controller/model/request"
	"crud/src/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateMembroController(c *gin.Context) {
	logger.Info("Init CreateMembro Controller",
		zap.String("journey", "createMembro"),
	)

	var membroRequest request.MembroRequest

	if err := c.ShouldBindJSON(&membroRequest); err != nil {
		logger.Error("Erro ao tentar validar info de membro", err,
			zap.String("journey", "createMembro"),
		)
		restErr := validation.ValidadeMembroError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewMembroDomain(
		membroRequest.Email,
		membroRequest.Password,
		membroRequest.Name,
		membroRequest.Age,
	)

	if err := domain.CreateMembroModel(); err != nil {
		c.JSON(err.Code, membroRequest)
		return
	}

	logger.Info("Membro criado com sucesso",
		zap.String("journey", "createMembro"))

	c.JSON(201, "")
}
