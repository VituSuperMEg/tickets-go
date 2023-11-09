package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base     `valid:"required"`
	Login    string `json:"login" gorm:"type:varchar(200)" valid:"notnull"`
	Email    string `json:"email" gorm:"type:varchar(200)" valid:"notnull"`
	Password string `json:"password" gorm:"type:varchar(200)" valid:"notnull"`
	Perfil   string `json:"perfil" gorm:"type:varchar(200)" valid:"notnull"`
}

type UserRepositoryInterface interface {
	Register(user *User) error
	Save(user *User) error
	Find(id string) (*User, error)
	Delete(id string) error
}

func (user *User) IsValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}
func (User) TableName() string {
	return "users"
}

func NewUser(login string, email string, password string, perfil string) (*User, error) {
	user := User{
		Login:    login,
		Email:    email,
		Password: password,
		Perfil:   perfil,
	}
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.IsValid()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
