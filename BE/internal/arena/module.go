package arena

import (
	"scout-arena/internal/arena/handler"
	"scout-arena/internal/arena/repository"
	"scout-arena/internal/arena/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitModule(router *gin.RouterGroup, database *gorm.DB) {
	seasonRepository := repository.NewSeasonRepository(database)
	seasonService := service.NewSeasonService(seasonRepository)
	seasonHandler := handler.NewSeasonHandler(seasonService)

	roundRepository := repository.NewRoundRepository(database)
	roundService := service.NewRoundService(roundRepository)
	roundHandler := handler.NewRoundHandler(roundService)

	challengeRepository := repository.NewChallengeRepository(database)
	challengeService := service.NewChallengeService(challengeRepository)
	challengeHandler := handler.NewChallengeHandler(challengeService)

	seasons := router.Group("/seasons")
	{
		seasons.POST("", seasonHandler.Create)
		seasons.PUT("/:id", seasonHandler.Update)
		seasons.DELETE("/:id", seasonHandler.Delete)
		seasons.GET("/:id", seasonHandler.GetById)
		seasons.GET("/:id/rounds", roundHandler.GetBySeasonId)
	}

	rounds := router.Group("/rounds")
	{
		rounds.POST("", roundHandler.Create)
		rounds.PUT("/:id", roundHandler.Update)
		rounds.DELETE("/:id", roundHandler.Delete)
		rounds.GET("/:id", roundHandler.GetById)
		rounds.GET("/:id/challenges", challengeHandler.GetByRoundId)
	}

	challenges := router.Group("/challenges")
	{
		challenges.POST("", challengeHandler.Create)
		challenges.PUT("/:id", challengeHandler.Update)
		challenges.DELETE("/:id", challengeHandler.Delete)
		challenges.GET("/:id", challengeHandler.GetById)
	}
}
