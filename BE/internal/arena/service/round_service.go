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
	repository RoundRepository
}

func NewRoundService(repo RoundRepository) *roundService {
	return &roundService{repository: repo}
}

func (service *roundService) Create(ctx context.Context, request *model.RoundRequest) (*model.RoundResponse, error) {
	var newRound = request.ToEntity()

	createdRound, err := service.repository.Create(ctx, newRound)
	if err != nil {
		return nil, err
	}

	return service.GetById(ctx, createdRound.ID)
}

func (service *roundService) Update(ctx context.Context, request *model.RoundRequest, id uint) (*model.RoundResponse, error) {
	round, err := service.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	round.UpdateFromRequest(request)

	updatedRound, err := service.repository.Update(ctx, round)
	if err != nil {
		return nil, err
	}

	return service.GetById(ctx, updatedRound.ID)
}

func (service *roundService) Delete(ctx context.Context, id uint) error {
	if err := service.repository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (service *roundService) GetById(ctx context.Context, id uint) (*model.RoundResponse, error) {
	round, err := service.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return round.ToResponse(), nil
}

func (service *roundService) GetBySeasonId(ctx context.Context, seasonId uint) ([]*model.RoundResponse, error) {
	rounds, err := service.repository.GetBySeasonID(ctx, seasonId)
	if err != nil {
		return nil, err
	}

	var roundResponses []*model.RoundResponse
	for _, round := range rounds {
		roundResponses = append(roundResponses, round.ToResponse())
	}

	return roundResponses, nil
}
