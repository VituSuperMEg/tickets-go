package repository

import (
	"fmt"

	"github.com/VituSuperMEg/tickets-go/domain/model"
	"github.com/jinzhu/gorm"
)

type FilmRepositoryDB struct {
	DB *gorm.DB
}

func (f *FilmRepositoryDB) Register(film *model.Film) error {
	err := f.DB.Create(film).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *FilmRepositoryDB) Save(film *model.Film) error {
	err := f.DB.Save(film).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *FilmRepositoryDB) Find(id string) (*model.Film, error) {
	var film model.Film
	err := f.DB.First(&film, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("no film was found: %v", err)
	}
	return &film, nil
}

func (f *FilmRepositoryDB) Delete(id string) error {
	err := f.DB.Delete(&model.Film{}, "id = ?", id).Error
	if err != nil {
		return fmt.Errorf("failed to delete film: %v", err)
	}
	return nil
}
func (f *FilmRepositoryDB) List() ([]*model.Film, error) {
	var films []*model.Film
	err := f.DB.Find(&films).Error
	if err != nil {
		return nil, err
	}
	return films, nil
}
