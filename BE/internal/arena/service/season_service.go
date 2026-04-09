package service

import (
	"context"
	"scout-arena/internal/arena/model"
)

type SeasonRepository interface {
	Create(ctx context.Context, season *model.Season) (*model.Season, error)
	Update(ctx context.Context, season *model.Season) (*model.Season, error)
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*model.Season, error)
}

type seasonService struct {
	repo SeasonRepository
}

// trong go khi dùng construction để khởi tạo sẽ return con trỏ, tránh việc copy các resource lớn
func NewSeasonService(repo SeasonRepository) *seasonService {
	return &seasonService{repo: repo}
}

func (s *seasonService) Create(ctx context.Context, request *model.SeasonRequest) (*model.SeasonResponse, error) {
	var newSeason = request.ToEntity()

	createdSeason, err := s.repo.Create(ctx, newSeason)

	if err != nil {
		return nil, err
	}

	return createdSeason.ToResponse(), nil
}
