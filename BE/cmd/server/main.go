package main

import (
	"log"
	"os"
	"scout-arena/internal/arena"
	"scout-arena/internal/db"
	"scout-arena/internal/validator"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Lỗi khi tải file .env: %v", err)
	}

	dbCfg := db.Config{
		Driver: os.Getenv("DB_DRIVER"),
		DSN:    os.Getenv("DB_DSN"),
	}

	database, err := db.NewDatabase(dbCfg)
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	conn := database.GetDB()

	if err := db.MigrateDB(conn); err != nil {
		log.Fatalf("Lỗi khi migrate database: %v", err)
	}
	log.Println("Database migrated successfully")

	r := gin.Default()
	r.SetTrustedProxies(nil)
	validator.Init()

	apiV1 := r.Group("/api/v1")

	arena.InitModule(apiV1, conn)
	// user.InitModule(apiV1, conn)

	port := os.Getenv("APP_PORT")

	r.Run(port)
}
