package repository

import (
	"fmt"

	"github.com/VituSuperMEg/tickets-go/domain/model"
	"github.com/jinzhu/gorm"
)

type UserRepositoryDB struct {
	DB *gorm.DB
}

func (u *UserRepositoryDB) Register(user *model.User) error {
	err := u.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *UserRepositoryDB) Save(user *model.User) error {
	err := u.DB.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *UserRepositoryDB) Find(id string) (*model.User, error) {
	var user model.User
	u.DB.First(&user, "id = ?", id)
	if user.ID == "" {
		return nil, fmt.Errorf("no ticket was found")
	}
	return &user, nil
}
func (u *UserRepositoryDB) Delete(id string) error {
	err := u.DB.Delete(&model.User{}, "id = ?", id).Error
	if err != nil {
		return fmt.Errorf("failed to delete film: %v", err)
	}
	return nil
}
