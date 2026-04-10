package arena

import (
	"scout-arena/internal/arena/handler"
	"scout-arena/internal/arena/repository"
	"scout-arena/internal/arena/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitModule(router *gin.RouterGroup, database *gorm.DB) {
	repository := repository.NewSeasonRepository(database)
	service := service.NewSeasonService(repository)
	handler := handler.NewSeasonHandler(service)

	seasons := router.Group("/seasons")
	{
		seasons.POST("", handler.Create)
		seasons.PUT("/:id", handler.Update)
		seasons.DELETE("/:id", handler.Delete)
		seasons.GET("/:id", handler.GetById)
	}
}
