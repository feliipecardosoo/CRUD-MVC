package view

import (
	"crud/src/controller/model/response"
	"crud/src/model"
)

func ConvertMembroDomainToResponse(membroDomain model.MembroDomainInterface) response.MembroResponse {
	return response.MembroResponse{
		Name:       membroDomain.GetName(),
		Status:     membroDomain.GetStatus(),
		DataStatus: membroDomain.GetDataStatus(),
	}
}

func ConvertAdminDomainToResponse(adminDomain model.AdminDomainInterface) response.AdminResponse {
	return response.AdminResponse{
		Usuario: adminDomain.GetUsuario(),
	}
}
