package repository

import (
	"context"
	"scout-arena/internal/arena/model"

	"gorm.io/gorm"
)

type SeasonRepository interface {
	Create(ctx context.Context, season *model.Season) (*model.Season, error)
}

type seasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) SeasonRepository {
	return &seasonRepository{db: db}
}

func (s *seasonRepository) Create(ctx context.Context, season *model.Season) (*model.Season, error) {
	err := s.db.WithContext(ctx).Create(season).Error
	if err != nil {
		return nil, err
	}
	return season, nil
}
