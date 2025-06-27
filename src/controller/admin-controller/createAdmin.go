package controller

import (
	"crud/src/configuration/logger"
	"crud/src/configuration/validation"
	"crud/src/controller/model/request"

	"crud/src/model"
	"crud/src/view"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *adminControllerInterface) CreateAdminController(c *gin.Context) {
	logger.Info("Init CreateMembro Controller",
		zap.String("journey", "createMembro"),
	)

	var adminRequest request.AdminRequest

	if err := BindJSONStrict(c, &adminRequest); err != nil {
		logger.Error("Erro ao tentar validar info do admin", err,
			zap.String("journey", "createAdmin"),
		)
		restErr := validation.BindAndValidateStrict(c, err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewAdminDomain(
		adminRequest.Usuario,
		adminRequest.Senha,
	)

	if err := uc.service.CreateAdminModel(domain); err != nil {
		c.JSON(err.Code, adminRequest)
		return
	}

	logger.Info("Membro criado com sucesso",
		zap.String("journey", "createMembro"),
		zap.String("name", domain.GetUsuario()),
	)

	c.JSON(201, view.ConvertAdminDomainToResponse(
		domain,
	))
}
