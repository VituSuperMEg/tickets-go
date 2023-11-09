package usecase

import (
	"errors"

	"github.com/VituSuperMEg/tickets-go/domain/model"
)

type TicketUseCast struct {
	TicketsRepository model.TicketsRepositoryInterface
}

func (t *TicketUseCast) Register(name string, status bool, films *model.Film) (*model.Tickets, error) {
	ticket, err := model.NewTickets(name, status, films)
	if err != nil {
		return nil, err
	}
	t.TicketsRepository.Save(ticket)

	if ticket.Base.ID != "" {
		return ticket, nil
	}

	return nil, errors.New("unable to process this transaction")

}
