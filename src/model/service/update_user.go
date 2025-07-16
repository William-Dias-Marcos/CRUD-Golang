package service

import (
	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/rest_err"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/model"
)

func (*userDomainService) UpdateUser(
	userId string, 
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}