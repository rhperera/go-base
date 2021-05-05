package domain

type User struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Mobile string   `json:"mobile"`
	Roles  []string `json:"roles"`
	Avatar string   `json:"avatar"`
	Admin  bool     `json:"admin"`
}

type UserService interface {
	GetByID(id int64) *User
}

type UserRepo interface {
	GetByID(id int64) (*User, error)
}
