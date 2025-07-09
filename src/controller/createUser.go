package controller

import (
	"github.com/William-Dias-Marcos/CRUD-Golang/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){

	err := rest_err.NewBadRequestError("Chamou errado")
	c.JSON(err.Code, err)
}
