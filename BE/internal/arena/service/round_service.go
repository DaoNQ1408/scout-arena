package service

import (
	"context"
	"scout-arena/internal/arena/model"
)

type RoundRepository interface {
	Create(ctx context.Context, round *model.Round) (*model.Round, error)
	Update(ctx context.Context, round *model.Round) (*model.Round, error)
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*model.Round, error)
	GetBySeasonID(ctx context.Context, seasonID uint) ([]*model.Round, error)
}
type roundService struct {
	repo RoundRepository
}

func NewRoundService(repo RoundRepository) *roundService {
	return &roundService{repo: repo}
}

func (s *roundService) Create(ctx context.Context, request *model.RoundRequest) (*model.RoundResponse, error) {
	var newRound = request.ToEntity()

	createdRound, err := s.repo.Create(ctx, newRound)
	if err != nil {
		return nil, err
	}

	return createdRound.ToResponse(), nil
}

func (s *roundService) Update(ctx context.Context, request *model.RoundRequest, id uint) (*model.RoundResponse, error) {
	round, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	round.UpdateFromRequest(request)

	updatedRound, err := s.repo.Update(ctx, round)
	if err != nil {
		return nil, err
	}

	return updatedRound.ToResponse(), nil
}

func (s *roundService) Delete(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *roundService) GetById(ctx context.Context, id uint) (*model.RoundResponse, error) {
	round, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return round.ToResponse(), nil
}

func (s *roundService) GetBySeasonId(ctx context.Context, seasonId uint) ([]*model.RoundResponse, error) {
	rounds, err := s.repo.GetBySeasonID(ctx, seasonId)
	if err != nil {
		return nil, err
	}

	var roundResponses []*model.RoundResponse
	for _, round := range rounds {
		roundResponses = append(roundResponses, round.ToResponse())
	}

	return roundResponses, nil
}
