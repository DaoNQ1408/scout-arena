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

// Create GoDoc
// @Summary      Tạo mới một vòng đấu
// @Description  Tạo vòng đấu mới với tên, hình ảnh và thời gian bắt đầu/kết thúc
// @Tags         rounds
// @Accept       json
// @Produce      json
// @Param        season  body      model.RoundRequest  true  "Thông tin vòng đấu"
// @Success      201     {object}  model.RoundResponse
// @Failure      400     {object}  map[string]string "Lỗi validation"
// @Router       /rounds [post]
func (handler *roundHandler) Create(c *gin.Context) {
	var request model.RoundRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.service.Create(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Update godoc
// @Summary      Cập nhật vòng đấu
// @Description  Cập nhật thông tin chi tiết của một vòng đấu dựa trên ID
// @Tags         rounds
// @Accept       json
// @Produce      json
// @Param        id      path      int                  true  "Round ID"
// @Param        season  body      model.RoundRequest  true  "Dữ liệu cập nhật"
// @Success      200     {object}  model.RoundResponse
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Router       /rounds/{id} [put]
func (handler *roundHandler) Update(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	var request model.RoundRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.service.Update(c.Request.Context(), &request, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete godoc
// @Summary      Xóa vòng đấu
// @Description  Xóa (soft delete) dựa trên ID
// @Tags         rounds
// @Produce      json
// @Param        id   path      int  true  "Round ID"
// @Success      200  {object}  map[string]string "message: Round deleted successfully"
// @Failure      404  {object}  map[string]string
// @Router       /rounds/{id} [delete]
func (handler *roundHandler) Delete(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	if err := handler.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Round deleted successfully"})
}

// GetById godoc
// @Summary      Lấy chi tiết vòng đấu
// @Tags         rounds
// @Produce      json
// @Param        id   path      int  true  "Round ID"
// @Success      200  {object}  model.RoundResponse
// @Failure      404  {object}  map[string]string "Không tìm thấy"
// @Router       /rounds/{id} [get]
func (handler *roundHandler) GetById(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	response, err := handler.service.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetBySeasonId godoc
// @Summary      Lấy chi tiết vòng đấu theo mùa giải
// @Tags         rounds
// @Produce      json
// @Param        id   path      int  true  "Round ID"
// @Success      200  {object}  model.RoundResponse
// @Failure      404  {object}  map[string]string "Không tìm thấy"
// @Router       /rounds/{id} [get]
func (handler *roundHandler) GetBySeasonId(c *gin.Context) {
	seasonId, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	responses, err := handler.service.GetBySeasonId(c.Request.Context(), seasonId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}
