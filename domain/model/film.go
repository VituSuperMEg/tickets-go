package model

type Film struct {
	Base       `valid:"required"`
	Film_name  string `json:"session"`
	Film_count int64  `json:"film_count"`
	Film_time  int64  `json:"film_time"`
}
