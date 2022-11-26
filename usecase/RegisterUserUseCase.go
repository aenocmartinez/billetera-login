package usecase

import (
	"billetera-login/domain"
	"billetera-login/infraestructure/repository"
	"errors"
)

type RegisterUserUseCase struct {
}

func (useCase RegisterUserUseCase) Execute(userDto UserDTO) error {

	repository := repository.NewMySQL()

	user := repository.FindUserByEmail(userDto.Email)

	if user.Exists() {
		return errors.New("el usuario ya existe")
	}

	user = *domain.NewUser(userDto.Name, userDto.Email, userDto.Cellphone, userDto.Password, userDto.Role)
	user.SetRepository(repository)

	return user.Create()
}
