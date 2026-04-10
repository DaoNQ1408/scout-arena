package repository

import (
	"context"
	"scout-arena/internal/arena/model"

	"gorm.io/gorm"
)

type roundRepository struct {
	db *gorm.DB
}

func NewRoundRepository(db *gorm.DB) *roundRepository {
	return &roundRepository{db: db}
}

func (r *roundRepository) Create(ctx context.Context, round *model.Round) (*model.Round, error) {
	err := r.db.WithContext(ctx).Create(round).Error
	if err != nil {
		return nil, err
	}

	return round, nil
}

func (r *roundRepository) Update(ctx context.Context, round *model.Round) (*model.Round, error) {
	err := r.db.WithContext(ctx).Save(round).Error
	if err != nil {
		return nil, err
	}

	return round, nil
}

func (r *roundRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.WithContext(ctx).Delete(&model.Round{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roundRepository) GetById(ctx context.Context, id uint) (*model.Round, error) {
	var round model.Round

	if err := r.db.WithContext(ctx).First(&round, id).Error; err != nil {
		return nil, err
	}

	return &round, nil
}

func (r *roundRepository) GetBySeasonID(ctx context.Context, seasonID uint) ([]*model.Round, error) {
	var rounds []*model.Round

	err := r.db.WithContext(ctx).Where("season_id = ?", seasonID).Find(&rounds).Error
	if err != nil {
		return nil, err
	}

	return rounds, nil
}
