package domain

type RegisterUserRepository interface {
	FindUserByEmail(email string) User
	CreateUser(user User) error
}
