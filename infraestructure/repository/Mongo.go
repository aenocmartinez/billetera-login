package repository

import (
	"billetera-login/domain"
	"errors"
	"fmt"
)

type MongoDB struct{}

func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (mongo *MongoDB) FindUserByEmail(email string) domain.User {
	fmt.Println("Buscando usuario en Mongo")
	return domain.User{}
}

func (mongo *MongoDB) CreateUser(user domain.User) error {
	return errors.New("")
}
