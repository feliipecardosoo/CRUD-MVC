package view

import (
	"crud/src/controller/model/response"
	"crud/src/model"
)

func ConvertDomainToResponse(membroDomain model.MembroDomainInterface) response.MembroResponse {
	return response.MembroResponse{
		Id:    "",
		Email: membroDomain.GetEmail(),
		Name:  membroDomain.GetName(),
		Age:   membroDomain.GetAge(),
	}
}
