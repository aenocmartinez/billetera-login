package controller

import (
	formrequest "billetera-login/infraestructure/controller/form-request"
	"billetera-login/usecase"
	"errors"

	"github.com/gin-gonic/gin"
)

type RegisterUserController struct {
}

func NewRegisterUserController() *RegisterUserController {
	return &RegisterUserController{}
}

func (controller *RegisterUserController) RegisterUser(c *gin.Context) error {
	var err error
	var req formrequest.RegisterUserFormRequests
	err = c.ShouldBindJSON(&req)
	if err != nil {
		return err
	}

	if req.Role != "Consumidor" && req.Role != "Comercio" {
		return errors.New("roles no válidos")
	}

	useCase := usecase.RegisterUserUseCase{}

	return useCase.Execute(usecase.UserDTO{
		Name:      req.Name,
		Cellphone: req.Cellphone,
		Email:     req.Email,
		Role:      req.Role,
		Password:  req.Password,
	})
}
