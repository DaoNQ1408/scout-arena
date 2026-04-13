package repository

import (
	"context"
	"scout-arena/internal/arena/model"

	"gorm.io/gorm"
)

type seasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) *seasonRepository {
	return &seasonRepository{db: db}
}

func (repository *seasonRepository) Create(ctx context.Context, season *model.Season) (*model.Season, error) {
	err := repository.db.WithContext(ctx).Create(season).Error
	if err != nil {
		return nil, err
	}

	return season, nil
}

func (repository *seasonRepository) Update(ctx context.Context, season *model.Season) (*model.Season, error) {
	err := repository.db.WithContext(ctx).Save(season).Error
	if err != nil {
		return nil, err
	}

	return season, nil
}

func (repository *seasonRepository) Delete(ctx context.Context, id uint) error {
	err := repository.db.WithContext(ctx).Delete(&model.Season{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *seasonRepository) GetById(ctx context.Context, id uint) (*model.Season, error) {
	var season model.Season

	if err := repository.db.WithContext(ctx).First(&season, id).Error; err != nil {
		return nil, err
	}

	return &season, nil
}
