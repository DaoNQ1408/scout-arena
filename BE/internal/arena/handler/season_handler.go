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

// Create GoDoc
// @Summary      Tạo mới một mùa giải
// @Description  Tạo mùa giải mới với tên, hình ảnh và thời gian bắt đầu/kết thúc
// @Tags         seasons
// @Accept       json
// @Produce      json
// @Param        season  body      model.SeasonRequest  true  "Thông tin mùa giải"
// @Success      201     {object}  model.SeasonResponse
// @Failure      400     {object}  map[string]string "Lỗi validation"
// @Router       /seasons [post]
func (handler *seasonHandler) Create(c *gin.Context) {
	var request model.SeasonRequest

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
// @Summary      Cập nhật mùa giải
// @Description  Cập nhật thông tin chi tiết của một mùa giải dựa trên ID
// @Tags         seasons
// @Accept       json
// @Produce      json
// @Param        id      path      int                  true  "Season ID"
// @Param        season  body      model.SeasonRequest  true  "Dữ liệu cập nhật"
// @Success      200     {object}  model.SeasonResponse
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Router       /seasons/{id} [put]
func (handler *seasonHandler) Update(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	var request model.SeasonRequest

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
// @Summary      Xóa mùa giải
// @Description  Xóa (soft delete) một mùa giải dựa trên ID
// @Tags         seasons
// @Produce      json
// @Param        id   path      int  true  "Season ID"
// @Success      200  {object}  map[string]string "message: Season deleted successfully"
// @Failure      404  {object}  map[string]string
// @Router       /seasons/{id} [delete]
func (handler *seasonHandler) Delete(c *gin.Context) {
	id, validParam := pkg.GetIDParam(c, "id")
	if !validParam {
		return
	}

	err := handler.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Season deleted successfully"})
}

// GetById godoc
// @Summary      Lấy chi tiết mùa giải
// @Description  Lấy thông tin chi tiết của một mùa giải dựa trên ID
// @Tags         seasons
// @Produce      json
// @Param        id   path      int  true  "Season ID"
// @Success      200  {object}  model.SeasonResponse
// @Failure      404  {object}  map[string]string "Không tìm thấy"
// @Router       /seasons/{id} [get]
func (handler *seasonHandler) GetById(c *gin.Context) {
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
