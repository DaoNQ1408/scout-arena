package database

import (
	arenaModel "scout-arena/internal/arena/model"
	participationModel "scout-arena/internal/participation/model"
	userModel "scout-arena/internal/user/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&userModel.Rank{},
		&userModel.Team{},
		&userModel.Role{},
		&userModel.User{},

		&arenaModel.Season{},
		&arenaModel.Round{},
		&arenaModel.Challenge{},

		&participationModel.UserChallengeRecord{},
		&participationModel.UserRoundProgress{},
		&participationModel.UserSeasonStat{},
	)
}
