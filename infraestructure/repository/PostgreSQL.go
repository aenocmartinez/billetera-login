package repository

import (
	"billetera-login/domain"
	"errors"
	"fmt"
)

type PostgreSQL struct{}

func NewPostgreSQL() *PostgreSQL {
	return &PostgreSQL{}
}

func (postgres *PostgreSQL) FindUserByEmail(email string) domain.User {
	fmt.Println("Buscando usuario en PostgreSQL")
	return domain.User{}
}

func (postgres *PostgreSQL) CreateUser(user domain.User) error {
	return errors.New("")
}
