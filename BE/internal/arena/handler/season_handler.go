package handler

import (
	"context"
	"net/http"
	"scout-arena/internal/arena/model"
	"scout-arena/internal/pkg"

	"github.com/gin-gonic/gin"
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

func (s *seasonHandler) Create(c *gin.Context) {
	var request model.SeasonRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := s.service.Create(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (s *seasonHandler) Update(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	var request model.SeasonRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := s.service.Update(c.Request.Context(), &request, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (s *seasonHandler) Delete(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	err := s.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Season deleted successfully"})
}

func (s *seasonHandler) GetById(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	response, err := s.service.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
