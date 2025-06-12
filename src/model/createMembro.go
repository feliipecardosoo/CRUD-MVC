package model

import (
	"crud/src/configuration/logger"
	resterr "crud/src/configuration/rest-err"

	"go.uber.org/zap"
)

func (ud *MembroDomain) CreateMembro() *resterr.RestErr {
	logger.Info("Init createMembro model", zap.String("jorney", "createMembro"))
	ud.EncryptPassword()
	return nil
}
