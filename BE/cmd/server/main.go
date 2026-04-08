package main

import (
	"log"
	"os"
	arenaModel "scout-arena/internal/arena/model"
	"scout-arena/internal/db"
	participantionModel "scout-arena/internal/participation/model"
	userModel "scout-arena/internal/user/model"
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

	err = conn.AutoMigrate(
		&userModel.Rank{},
		&userModel.Team{},
		&userModel.Role{},
		&userModel.User{},

		&arenaModel.Challenge{},
		&arenaModel.Round{},
		&arenaModel.Season{},

		&participantionModel.UserChallengeRecord{},
		&participantionModel.UserRoundProgress{},
		&participantionModel.UserSeasonStat{},
	)

	if err != nil {
		log.Printf("Lỗi khi migrate database: %v", err)
	} else {
		log.Println("Database migrated successfully")
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)
	validator.Init()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080" // port mặc định
	}
	r.Run(port)
}
