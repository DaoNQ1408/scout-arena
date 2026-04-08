package repository

import (
	"context"
	"scout-arena/internal/arena/model"

	"gorm.io/gorm"
)

type ChallengeRepository interface {
	Create(ctx context.Context, request *model.ChallengeRequest) (*model.ChallengeResponse, error)
}

// constructor
type challengeRepository struct {
	db *gorm.DB
}

// inject dependency
func NewChallengeRepository(db *gorm.DB) ChallengeRepository {
	return &challengeRepository{db: db}
}

func (c challengeRepository) Create(ctx context.Context, request *model.ChallengeRequest) (*model.ChallengeResponse, error) {
	panic("implement me")
}
