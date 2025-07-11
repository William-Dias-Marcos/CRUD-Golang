package controller

import (
	"fmt"

	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/validation"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/controller/model/request"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){

	var userRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
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

	c.JSON(200, response)
}
