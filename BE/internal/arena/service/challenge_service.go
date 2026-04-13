package service

import (
	"context"
	"scout-arena/internal/arena/model"
)

type ChallengeRepository interface {
	Create(ctx context.Context, challenge *model.Challenge) (*model.Challenge, error)
	Update(ctx context.Context, challenge *model.Challenge) (*model.Challenge, error)
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*model.Challenge, error)
	GetByRoundId(ctx context.Context, roundId uint) ([]*model.Challenge, error)
}

type challengeService struct {
	repository ChallengeRepository
}

func NewChallengeService(repository ChallengeRepository) *challengeService {
	return &challengeService{repository: repository}
}

func (service *challengeService) Create(ctx context.Context, request *model.ChallengeRequest) (*model.ChallengeResponse, error) {
	var newChallenge = request.ToEntity()

	createdChallenge, err := service.repository.Create(ctx, newChallenge)
	if err != nil {
		return nil, err
	}

	return createdChallenge.ToResponse(), nil
}

func (service *challengeService) Update(ctx context.Context, request *model.ChallengeRequest, id uint) (*model.ChallengeResponse, error) {
	challenge, err := service.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	challenge.UpdateFromRequest(request)

	updatedChallenge, err := service.repository.Update(ctx, challenge)
	if err != nil {
		return nil, err
	}

	return updatedChallenge.ToResponse(), nil
}

func (service *challengeService) Delete(ctx context.Context, id uint) error {
	if err := service.repository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (service *challengeService) GetById(ctx context.Context, id uint) (*model.ChallengeResponse, error) {
	challenge, err := service.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return challenge.ToResponse(), nil
}

func (service *challengeService) GetByRoundId(ctx context.Context, roundId uint) ([]*model.ChallengeResponse, error) {
	challenges, err := service.repository.GetByRoundId(ctx, roundId)
	if err != nil {
		return nil, err
	}

	var responses []*model.ChallengeResponse
	for _, challenge := range challenges {
		responses = append(responses, challenge.ToResponse())
	}

	return responses, nil
}
