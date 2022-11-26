package repository

import (
	"billetera-login/domain"
	"errors"
	"fmt"
)

type MySQL struct{}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (mysql *MySQL) FindUserByEmail(email string) domain.User {
	fmt.Println("Buscando usuario en MySQL")
	return domain.User{}
}

func (mysql *MySQL) CreateUser(user domain.User) error {
	return errors.New("")
}
