package handler

import (
	"context"
	"net/http"
	"scout-arena/internal/arena/model"
	"scout-arena/internal/pkg"

	"github.com/gin-gonic/gin"
)

type RoundService interface {
	Create(ctx context.Context, request *model.RoundRequest) (*model.RoundResponse, error)
	Update(ctx context.Context, request *model.RoundRequest, id uint) (*model.RoundResponse, error)
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*model.RoundResponse, error)
	GetBySeasonId(ctx context.Context, seasonId uint) ([]*model.RoundResponse, error)
}

type roundHandler struct {
	service RoundService
}

func NewRoundHandler(service RoundService) *roundHandler {
	return &roundHandler{service: service}
}

func (r *roundHandler) Create(c *gin.Context) {
	var request model.RoundRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := r.service.Create(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (r *roundHandler) Update(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	var request model.RoundRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := r.service.Update(c.Request.Context(), &request, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (r *roundHandler) Delete(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	if err := r.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Round deleted successfully"})
}

func (r *roundHandler) GetById(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	response, err := r.service.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (r *roundHandler) GetBySeasonId(c *gin.Context) {
	seasonId, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	responses, err := r.service.GetBySeasonId(c.Request.Context(), seasonId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}
