package usecase

import (
	"errors"

	"github.com/VituSuperMEg/tickets-go/domain/model"
)

type FilmUseCast struct {
	FilmRepository model.FilmRepositoryInterface
}

func (f *FilmUseCast) Register(name string, count int, time int64, description string, image string) (*model.Film, error) {
	film, err := model.NewFilm(name, count, time, description, image)
	if err != nil {
		return nil, err
	}
	f.FilmRepository.Save(film)

	if film.Base.ID != "" {
		return film, nil
	}
	return nil, errors.New("unable to find film this repository describes")
}
func (f *FilmUseCast) List() ([]*model.Film, error) {
	films, err := f.FilmRepository.List()
	if err != nil {
		return nil, err
	}
	return films, nil
}
func (f *FilmUseCast) Find(id string) (*model.Film, error) {
	films, err := f.FilmRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return films, nil
}
