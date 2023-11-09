package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Film struct {
	Base        `valid:"required"`
	Film_name   string `json:"session" gorm:"type:varchar(255);column:session" valid:"notnull"`
	Film_count  int    `json:"film_count" gorm:"type:integer"  valid:"notnull"`
	Film_time   int64  `json:"film_time" gorm:"type:integer"  valid:"notnull"`
	Description string `json:"description" gorm:"type:varchar(255)"  valid:"notnull"`
	ImagePath   string `json:"image_path" gorm:"type:varchar(255)"  valid:"notnull"`
}
type FilmRepositoryInterface interface {
	Register(film *Film) error
	Save(film *Film) error
	List() ([]*Film, error)
	Find(id string) (*Film, error)
	Delete(id string) error
}

func (Film) TableName() string {
	return "session"
}
func (film *Film) IsValid() error {
	_, err := govalidator.ValidateStruct(film)
	if err != nil {
		return err
	}
	return nil
}
func List() ([]*Film, error) {
	var films []*Film
	return films, nil
}

func NewFilm(name string, count int, time_hour int64, description string, image string) (*Film, error) {
	film := Film{
		Film_name:   name,
		Film_count:  count,
		Film_time:   time_hour,
		Description: description,
		ImagePath:   image,
	}
	film.ID = uuid.NewV4().String()
	film.CreatedAt = time.Now()

	err := film.IsValid()
	if err != nil {
		return nil, err
	}

	return &film, nil
}
