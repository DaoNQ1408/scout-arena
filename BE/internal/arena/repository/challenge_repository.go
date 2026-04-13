package repository

import (
	"context"
	"scout-arena/internal/arena/model"

	"gorm.io/gorm"
)

// constructor
type challengeRepository struct {
	db *gorm.DB
}

// inject dependency
func NewChallengeRepository(db *gorm.DB) *challengeRepository {
	return &challengeRepository{db: db}
}

func (repository *challengeRepository) Create(ctx context.Context, challenge *model.Challenge) (*model.Challenge, error) {
	err := repository.db.WithContext(ctx).Create(challenge).Error
	if err != nil {
		return nil, err
	}

	return challenge, nil
}

func (repository *challengeRepository) Update(ctx context.Context, challenge *model.Challenge) (*model.Challenge, error) {
	err := repository.db.WithContext(ctx).Save(challenge).Error
	if err != nil {
		return nil, err
	}

	return challenge, nil
}

func (repository *challengeRepository) Delete(ctx context.Context, id uint) error {
	err := repository.db.WithContext(ctx).Delete(&model.Challenge{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *challengeRepository) GetById(ctx context.Context, id uint) (*model.Challenge, error) {
	var challenge model.Challenge

	if err := repository.db.WithContext(ctx).First(&challenge, id).Error; err != nil {
		return nil, err
	}

	return &challenge, nil
}

func (repository *challengeRepository) GetByRoundId(ctx context.Context, roundId uint) ([]*model.Challenge, error) {
	var challenges []*model.Challenge

	err := repository.db.WithContext(ctx).Where("round_id = ?", roundId).Find(&challenges).Error
	if err != nil {
		return nil, err
	}

	var responses []*model.Challenge
	for _, challenge := range challenges {
		responses = append(responses, challenge)
	}

	return responses, nil
}
