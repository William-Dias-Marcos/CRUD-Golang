package controller

import (
	"net/http"

	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/logger"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/validation"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/controller/model/request"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/model"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context){
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err !=nil {
		c.JSON(err.Code, err)
		return
	}


	logger.Info("CreateUser controller executed successfully",
		zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
