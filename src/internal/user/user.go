package user

type User struct {
	id       int64
	email    string
	password string
	token    string
}

func NewUser(id int64, email string, password string) *User {
	return &User{id: id, email: email, password: password}
}

func (u *User) Id() int64 {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) Token() string {
	return u.token
}

func (u *User) SetToken(token string) {
	u.token = token
}
