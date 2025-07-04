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

func (uc *membroControllerInterface) CreateMembroController(c *gin.Context) {
	logger.Info("Init CreateMembro Controller",
		zap.String("journey", "createMembro"),
	)

	var membroRequest request.MembroRequest

	if err := BindJSONStrict(c, &membroRequest); err != nil {
		logger.Error("Erro ao tentar validar info de membro", err,
			zap.String("journey", "createMembro"),
		)
		restErr := validation.BindAndValidateStrict(c, err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewMembroDomain(
		membroRequest.Name,
		membroRequest.DataNascimento,
		membroRequest.AnoBatismo,
		membroRequest.Sexo,
		membroRequest.EstadoCivil,
		membroRequest.DataCasamento,
		membroRequest.NomeConjuge,
		membroRequest.Filho,
		membroRequest.Email,
		membroRequest.Telefone,
		model.ConvertEnderecoRequest(membroRequest.Endereco),
	)

	if err := uc.service.CreateMembroModel(domain); err != nil {
		c.JSON(err.Code, membroRequest)
		return
	}

	logger.Info("Membro criado com sucesso",
		zap.String("journey", "createMembro"),
		zap.String("name", domain.GetName()),
	)

	c.JSON(201, view.ConvertMembroDomainToResponse(
		domain,
	))
}
