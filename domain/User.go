package domain

type User struct {
	Id         int64
	Name       string
	Email      string
	Cellphone  string
	Password   string
	Role       string
	State      bool
	repository RegisterUserRepository
}

func NewUser(name, email, cellphone, password, role string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Cellphone: cellphone,
		Password:  password,
		Role:      role,
	}
}

func (u *User) SetRepository(repository RegisterUserRepository) {
	u.repository = repository
}

func (u *User) Create() error {
	return u.repository.CreateUser(*u)
}

func (u *User) Update() bool {
	return false
}

func (u *User) Exists() bool {
	return u.Id > 0
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
