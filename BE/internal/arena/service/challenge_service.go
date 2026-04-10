package service

import (
	"context"
	"scout-arena/internal/arena/model"
)

type ChallengeRepository interface {
	Create(ctx context.Context, request *model.Challenge) (*model.Challenge, error)
	Update(ctx context.Context, request *model.Challenge) (*model.Challenge, error)
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*model.Challenge, error)
}
