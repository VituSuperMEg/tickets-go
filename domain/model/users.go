package model

type User struct {
	Base     `valid:"required"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newUser(name string, email string, password string) (*User, error) {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return &user, nil
}
