package controller

import (
	"crud/src/configuration/logger"
	"crud/src/configuration/validation"
	"crud/src/controller/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateMembro(c *gin.Context) {
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

	c.JSON(201, membroRequest)
}
