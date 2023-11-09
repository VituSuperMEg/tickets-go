package usecase

import (
	"errors"

	"github.com/VituSuperMEg/tickets-go/domain/model"
)

type FilmUseCast struct {
	FilmRepository model.FilmRepositoryInterface
}

func (f *FilmUseCast) Register(name string, count int, time int64) (*model.Film, error) {
	film, err := model.NewFilm(name, count, time)
	if err != nil {
		return nil, err
	}
	f.FilmRepository.Save(film)

	if film.Base.ID != "" {
		return film, nil
	}
	return nil, errors.New("unable to find film this repository describes")
}
