package tickets

import (
	"github.com/Santiagolaiton27/desafio-go-web-santiago-laiton/internal/domain"
)

type Service interface {
	GetTotalTickets(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (float64, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(destination string) (tickets []domain.Ticket, err error) {
	tickets, err = s.repository.GetTicketByDestination(destination)
	if err != nil {
		return tickets, err
	}
	return tickets, err
}

func (s *service) AverageDestination(destination string) (float64, error) {
	return 0.0, nil
}
