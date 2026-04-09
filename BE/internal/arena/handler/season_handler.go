package handler

import (
	"context"
	"net/http"
	"scout-arena/internal/arena/model"

	"github.com/gin-gonic/gin"
)

type SeasonService interface {
	Create(ctx context.Context, request *model.SeasonRequest) (*model.SeasonResponse, error)
	//Update(ctx context.Context, request *model.SeasonRequest, id uint) (*model.SeasonResponse, error)
	//Delete(ctx context.Context, id uint) error
	//GetById(ctx context.Context, id uint) (*model.SeasonResponse, error)
}

type seasonHandler struct {
	service SeasonService
}

func NewSeasonHandler(service SeasonService) *seasonHandler {
	return &seasonHandler{service: service}
}

func (s *seasonHandler) Create(c *gin.Context) {
	var req model.SeasonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.service.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
