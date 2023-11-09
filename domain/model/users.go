package model

type User struct {
	Base     `valid:"required"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type TransactionsRepositoryInterface interface {
	Register(transaction *User) error
	Save(transaction *User) error
	Find(id string) (*User, error)
}

func newUser(name string, email string, password string) (*User, error) {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return &user, nil
}
