package model

type Tickets struct {
	Base         `valid:"required"`
	Session_name string  `json:"session"`
	Status       bool    `json:"status"`
	Films        []*Film `json:"films"`
}
