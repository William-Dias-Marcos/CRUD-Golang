package model

import (
	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/logger"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) CreateUser(UserDomain) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	return nil
}