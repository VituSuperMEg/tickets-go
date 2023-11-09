package model

type Tickets struct {
	Base         `valid:"required"`
	Session_name string  `json:"session"`
	Status       bool    `json:"status"`
	Films        []*Film `json:"films"`
}
type TicketsRepositoryInterface interface {
	Register(film *Film) error
	Save(ticket *Tickets) error
	Find(id string) (*Film, error)
	RegisterFilmeAndTicket(film []*Film) (*Film, error)
}

func NewTickets(session_name string, status bool, film *Film) (*Tickets, error) {
	tickets := Tickets{
		Session_name: session_name,
		Films:        []*Film{film},
		Status:       status,
	}
	return &tickets, nil
}
