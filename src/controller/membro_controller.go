package controller

import (
	resterr "crud/src/configuration/rest-err"
	"crud/src/model/service"
	"encoding/json"

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

// BindJSONStrict lê o JSON do request, rejeita campos extras e faz bind no obj.
func BindJSONStrict(c *gin.Context, obj interface{}) *resterr.RestErr {
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields() // rejeita campos extras
	if err := decoder.Decode(obj); err != nil {
		return resterr.NewBadRequestError("JSON inválido ou campos extras não permitidos: " + err.Error())
	}

	// Verifica se há dados extras após o JSON
	if decoder.More() {
		return resterr.NewBadRequestError("JSON contém dados extras após o objeto")
	}
	return nil
}
