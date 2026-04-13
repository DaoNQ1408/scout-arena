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

// Create GoDoc
// @Summary      Tạo mới một thử thách
// @Description  Tạo thử thách mới với tên, hình ảnh và thời gian bắt đầu/kết thúc
// @Tags         seasons
// @Accept       json
// @Produce      json
// @Param        season  body      model.ChallengeRequest  true  "Thông tin mùa giải"
// @Success      201     {object}  model.ChallengeResponse
// @Failure      400     {object}  map[string]string "Lỗi validation"
// @Router       /seasons [post]
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

// Update godoc
// @Summary      Cập nhật thử thách
// @Description  Cập nhật thông tin chi tiết của một thử thách dựa trên ID
// @Tags         challenges
// @Accept       json
// @Produce      json
// @Param        id      path      int                  true  "Round ID"
// @Param        challenge  body      model.ChallengeRequest  true  "Dữ liệu cập nhật"
// @Success      200     {object}  model.ChallengeResponse
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Router       /challenges/{id} [put]
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

// Delete godoc
// @Summary      Xóa thử thách
// @Description  Xóa (soft delete) dựa trên ID
// @Tags         challenges
// @Produce      json
// @Param        id   path      int  true  "Challenge ID"
// @Success      200  {object}  map[string]string "message: Challenge deleted successfully"
// @Failure      404  {object}  map[string]string
// @Router       /challenges/{id} [delete]
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

// GetById godoc
// @Summary      Lấy chi tiết thử thách
// @Tags         challenges
// @Produce      json
// @Param        id   path      int  true  "Challenge ID"
// @Success      200  {object}  model.ChallengeResponse
// @Failure      404  {object}  map[string]string "Không tìm thấy"
// @Router       /challenges/{id} [get]
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

// GetByRoundId godoc
// @Summary      Lấy chi tiết thử thách theo vòng đấu
// @Tags         challenges
// @Produce      json
// @Param        id   path      int  true  "Challenge ID"
// @Success      200  {object}  model.ChallengeResponse
// @Failure      404  {object}  map[string]string "Không tìm thấy"
// @Router       /challenges/{id} [get]
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
