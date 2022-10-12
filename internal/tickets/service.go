package tickets

import (
	"errors"

	"github.com/Santiagolaiton27/desafio-go-web-santiago-laiton/internal/domain"
)

type Service interface {
	GetTotalTickets(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (float64, error)
	GetAll() (tickets []domain.Ticket, err error)
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
	tickets, err := s.repository.GetAll()
	if err != nil {
		return 0.0, err
	}
	ticketsDestination, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0.0, err
	}

	if len(ticketsDestination) == 0 && len(tickets) == 0 {
		return 0.0, errors.New("No se encontraron tiquetes de esa ciudad de destino")
	}
	average := len(tickets) / len(ticketsDestination)
	return float64(average), nil
}

func (s *service) GetAll() (tickets []domain.Ticket, err error) {
	tickets, err = s.repository.GetAll()
	if err != nil {
		return tickets, err
	}
	return tickets, nil
}
