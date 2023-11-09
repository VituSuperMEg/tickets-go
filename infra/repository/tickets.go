package repository

import (
	"fmt"

	"github.com/VituSuperMEg/tickets-go/domain/model"
	"github.com/jinzhu/gorm"
)

type TicketsRepositoryDb struct {
	DB *gorm.DB
}

func (t *TicketsRepositoryDb) Register(ticket *model.Tickets) error {
	err := t.DB.Create(ticket).Error
	if err != nil {
		return err
	}
	return nil
}
func (t *TicketsRepositoryDb) Save(ticket *model.Tickets) error {
	err := t.DB.Create(ticket).Error
	if err != nil {
		return err
	}
	return nil
}
func (t *TicketsRepositoryDb) RegisterFilmeAndTicket(film *model.Film) error {
	err := t.DB.Create(film).Error
	if err != nil {
		return err
	}
	return nil
}
func (t *TicketsRepositoryDb) Find(id string) (*model.Tickets, error) {
	var ticket model.Tickets
	t.DB.First(&ticket, "id = ?", id)
	if ticket.ID == "" {
		return nil, fmt.Errorf("no ticket was found")
	}
	return &ticket, nil
}
