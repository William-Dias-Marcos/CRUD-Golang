package controller

import (
	"fmt"

	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/rest_err"
	"github.com/William-Dias-Marcos/CRUD-Golang/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){

	var userRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s\n" , err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
