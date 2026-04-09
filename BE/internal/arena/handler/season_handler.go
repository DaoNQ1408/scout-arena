package handler

import (
	"context"
	"scout-arena/internal/arena/model"
)

type SeasonService interface {
	Create(ctx context.Context, request *model.SeasonRequest) (*model.SeasonResponse, error)
	Update(ctx context.Context, request *model.SeasonRequest, id uint) (*model.SeasonResponse, error)
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*model.SeasonResponse, error)
}

type seasonHandler struct {
	service SeasonService
}

func NewSeasonHandler(service SeasonService) *seasonHandler {
	return &seasonHandler{service: service}
}
