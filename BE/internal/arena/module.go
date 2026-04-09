package arena

import (
	"scout-arena/internal/arena/handler"
	"scout-arena/internal/arena/repository"
	"scout-arena/internal/arena/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitModule(router *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewSeasonRepository(db)
	svc := service.NewSeasonService(repo)
	hdl := handler.NewSeasonHandler(svc)

	// 2. Khai báo routes cho module này
	seasons := router.Group("/seasons")
	{
		seasons.POST("", hdl.Create)
		// seasons.GET("/:id", hdl.GetByID)
	}
}
