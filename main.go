package main

import (
	"billetera-login/infraestructure/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Mundo GO",
		})
	})

	r.POST("/register-user", func(c *gin.Context) {

		ctr := controller.NewRegisterUserController()
		err := ctr.RegisterUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": "success"})
		}
	})

	r.Run()
}
