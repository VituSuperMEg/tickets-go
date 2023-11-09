package usecase

import (
	"errors"
	"fmt"

	"github.com/VituSuperMEg/tickets-go/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCast struct {
	UserRepository model.UserRepositoryInterface
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *UserUseCast) Register(name string, email string, password string, perfil string) (*model.User, error) {

	hash, err := hashPassword(password)

	if err != nil {
		fmt.Println("Erro a gerar senha!", err)
		return nil, err
	}
	user, err := model.NewUser(name, email, hash, perfil)

	if err != nil {
		return nil, err
	}
	u.UserRepository.Save(user)
	if user.Base.ID != "" {
		return user, nil
	}
	return nil, errors.New("unable to find film this repository describes")
}
