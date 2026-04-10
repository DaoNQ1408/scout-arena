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

func (s *seasonService) Update(ctx context.Context, request *model.SeasonRequest, id uint) (*model.SeasonResponse, error) {
	var season, err = s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	season.UpdateFromRequest(request)

	updatedSeason, err := s.repo.Update(ctx, season)
	if err != nil {
		return nil, err
	}

	return updatedSeason.ToResponse(), err
}

func (s *seasonService) Delete(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *seasonService) GetById(ctx context.Context, id uint) (*model.SeasonResponse, error) {
	season, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return season.ToResponse(), nil
}
