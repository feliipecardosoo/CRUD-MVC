package view

import (
	"crud/src/controller/model/response"
	"crud/src/model"
)

func ConvertDomainToResponse(membroDomain model.MembroDomainInterface) response.MembroResponse {
	return response.MembroResponse{
		Name:       membroDomain.GetName(),
		Status:     membroDomain.GetStatus(),
		DataStatus: membroDomain.GetDataStatus(),
		Validado:   membroDomain.GetValidado(),
	}
}
