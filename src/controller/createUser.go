package controller

import (
	"fmt"
	"net/http"

	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/logger"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/validation"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/controller/model/request"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	fmt.Println(userRequest)

	response := response.UserResponse{
		ID: "123456789",
		Name: userRequest.Name,
		Email: userRequest.Email,
		Age: userRequest.Age,
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
