package handler

import (
	"context"
	"net/http"
	"scout-arena/internal/arena/model"
	"scout-arena/internal/pkg"

	"github.com/gin-gonic/gin"
)

type ChallengeService interface {
	Create(ctx context.Context, request *model.ChallengeRequest) (*model.ChallengeResponse, error)
	Update(ctx context.Context, request *model.ChallengeRequest, id uint) (*model.ChallengeResponse, error)
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*model.ChallengeResponse, error)
	GetByRoundId(ctx context.Context, roundId uint) ([]*model.ChallengeResponse, error)
}

type challengeHandler struct {
	service ChallengeService
}

func NewChallengeHandler(service ChallengeService) *challengeHandler {
	return &challengeHandler{service: service}
}

func (handler *challengeHandler) Create(c *gin.Context) {
	var request model.ChallengeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.service.Create(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (handler *challengeHandler) Update(c *gin.Context) {
	id, validID := pkg.GetIDParam(c, "id")
	if !validID {
		return
	}

	var request model.ChallengeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.service.Update(c, &request, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *challengeHandler) Delete(c *gin.Context) {
	id, validID := pkg.GetIDParam(c, "id")
	if !validID {
		return
	}

	if err := handler.service.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Challenge deleted successfully"})
}

func (handler *challengeHandler) GetById(c *gin.Context) {
	id, validID := pkg.GetIDParam(c, "id")
	if !validID {
		return
	}

	response, err := handler.service.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *challengeHandler) GetByRoundId(c *gin.Context) {
	id, validID := pkg.GetIDParam(c, "id") // roundId
	if !validID {
		return
	}

	responses, err := handler.service.GetByRoundId(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}
