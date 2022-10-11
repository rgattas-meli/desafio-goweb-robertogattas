package tickets

import (
	"github.com/rgattas-meli/desafio-goweb-robertogattas/pkg/domain"
	"context"

)
type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(context.Context, string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string)	(int, error)

}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}


func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error)	{
	ps, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return ps, nil
}
func (s *service) GetTotalTickets(ctx context.Context,name  string) ([]domain.Ticket, error)	{
	ps, err := s.repository.GetTicketByDestination(ctx, name)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string)	(int, error)	{
	ps, err := s.repository.GetTicketsAvg(ctx , destination)
	if err != nil {
		return 0, err
	}

	return ps, nil
}