package main

import (
	"log"
	"os"
	_ "scout-arena/docs"
	"scout-arena/internal/arena"
	"scout-arena/internal/database"
	"scout-arena/internal/validator"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Scout Arena API
// @version 1.0
// @description API for managing seasons, rounds, and challenges in the Scout Arena application.
// @contact.name API Support
// @contact.url http://www.scoutarena.com/support
// @contact.email
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Lỗi khi tải file .env: %v", err)
	}

	databaseConfig := database.Config{
		Driver: os.Getenv("DB_DRIVER"),
		DSN:    os.Getenv("DB_DSN"),
	}

	newDatabase, err := database.NewDatabase(databaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	defer newDatabase.Close()

	connection := newDatabase.GetDB()

	if err := database.MigrateDB(connection); err != nil {
		log.Fatalf("Lỗi khi migrate database: %v", err)
	}
	log.Println("Database migrated successfully")

	r := gin.Default()
	r.SetTrustedProxies(nil)
	validator.Init()

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := r.Group("/api/v1")

	arena.InitModule(apiV1, connection)
	// user.InitModule(apiV1, conn)

	port := os.Getenv("APP_PORT")

	r.Run(port)
}
