package service

import (
	"crud/src/configuration/logger"
	resterr "crud/src/configuration/rest-err"
	"crud/src/model"

	"go.uber.org/zap"
)

func (ud *adminDomainService) CreateAdminModel(adminDomain model.AdminDomainInterface) *resterr.RestErr {
	logger.Info("Init createAdmin model", zap.String("jorney", "createAdmin"))
	return nil
}
