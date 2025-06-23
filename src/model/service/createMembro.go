package service

import (
	"crud/src/configuration/logger"
	resterr "crud/src/configuration/rest-err"
	"crud/src/model"
	"fmt"

	"go.uber.org/zap"
)

func (ud *membroDomainService) CreateMembroModel(userDomain model.MembroDomainInterface) *resterr.RestErr {
	logger.Info("Init createMembro model", zap.String("jorney", "createMembro"))
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetName())
	fmt.Println(userDomain.GetPassword())
	return nil
}
