package db

import (
	arenaModel "scout-arena/internal/arena/model"
	participationModel "scout-arena/internal/participation/model"
	userModel "scout-arena/internal/user/model"

	"gorm.io/gorm"
)

// MigrateDB thực hiện tự động tạo bảng cho toàn bộ hệ thống
func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		// User Group
		&userModel.Rank{},
		&userModel.Team{},
		&userModel.Role{},
		&userModel.User{},

		// Arena Group
		&arenaModel.Season{},
		&arenaModel.Round{},
		&arenaModel.Challenge{},

		// Participation Group
		&participationModel.UserChallengeRecord{},
		&participationModel.UserRoundProgress{},
		&participationModel.UserSeasonStat{},
	)
}
