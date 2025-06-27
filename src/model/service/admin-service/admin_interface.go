package service

import (
	resterr "crud/src/configuration/rest-err"
	"crud/src/model"
)

func NewAdminDomainService() AdminDomainService {
	return &adminDomainService{}
}

type adminDomainService struct {
}

type AdminDomainService interface {
	CreateAdminModel(model.AdminDomainInterface) *resterr.RestErr
}
