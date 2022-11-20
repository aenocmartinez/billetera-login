package domain

type User struct {
	Id        int64
	Name      string
	Email     string
	Cellphone string
	Password  string
	Role      string
	State     bool
}

func (u *User) NewUser(name, email, cellphone, password string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Cellphone: cellphone,
		Password:  password,
	}
}

func (u *User) Create() bool {
	return false
}

func (u *User) Update() bool {
	return false
}

func (u *User) Exists() bool {
	return false
}

func (u *User) Find(email string) User {
	return User{}
}

func (u *User) AssignRole(role string) bool {
	return false
}

func (u *User) ChangeState(state string) bool {
	return false
}
