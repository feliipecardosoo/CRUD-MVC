package service

import (
	resterr "crud/src/configuration/rest-err"
	"crud/src/model"
)

type membroDomainService struct {
}

type MembroDomainService interface {
	CreateMembroModel(model.MembroDomainInterface) *resterr.RestErr
	UpdateMembro(string, model.MembroDomainInterface) *resterr.RestErr
	FindMembro(string) (*model.MembroDomainInterface, *resterr.RestErr)
	DeleteMembro(string) *resterr.RestErr
}
