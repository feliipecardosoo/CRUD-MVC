package controller

import (
	resterr "crud/src/configuration/rest-err"
	service "crud/src/model/service/admin-service"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func NewAdminControllerInterface(serviceInterface service.AdminDomainService) AdminControllerInterface {
	return &adminControllerInterface{
		service: serviceInterface,
	}
}

type AdminControllerInterface interface {
	CreateAdminController(c *gin.Context)
}

type adminControllerInterface struct {
	service service.AdminDomainService
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
