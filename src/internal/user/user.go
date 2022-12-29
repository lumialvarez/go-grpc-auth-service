package user

type User struct {
	id       int64
	name     string
	userName string
	email    string
	password string
	token    string
	role     Role
	status   bool
}

type Role string

const (
	RolUser  Role = "role_user"
	RolAdmin      = "role_admin"
)

func NewUser(id int64, name string, userName string, email string, password string, token string, role Role, status bool) *User {
	return &User{id: id, name: name, userName: userName, email: email, password: password, token: token, role: role, status: status}
}

func (u *User) Id() int64 {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) UserName() string {
	return u.userName
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Token() string {
	return u.token
}

func (u *User) Role() Role {
	return u.role
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetToken(token string) {
	u.token = token
}

func (u *User) Status() bool {
	return u.status
}

func (u *User) SetStatus(status bool) {
	u.status = status
}
