package model

import (
	arenaModel "scout-arena/internal/arena/model"
	userModel "scout-arena/internal/user/model"

	"gorm.io/gorm"
)

type UserRoundProgress struct {
	gorm.Model
	TotalPoint uint                `json:"total_point"` // total till this round
	RoundPoint uint                `json:"round_point"` // point in this round
	Status     ParticipationStatus `json:"status" gorm:"type:varchar(20);default:'DRAFT';not null"`

	UserID uint           `json:"user_id" gorm:"uniqueIndex:idx_user_round"`
	User   userModel.User `json:"user" gorm:"foreignkey:UserID"`

	RoundID   uint             `json:"round_id" gorm:"uniqueIndex:idx_user_round"`
	Round     arenaModel.Round `json:"round" gorm:"foreignkey:RoundID"`
	DeletedAt gorm.DeletedAt   `gorm:"uniqueIndex:idx_user_round"`

	RankAtStartID uint           `json:"rank_at_start_id"`
	RankAtStart   userModel.Rank `json:"rank_at_start" gorm:"foreignkey:RankAtStartID"`
}

type UserChallengeRecord struct {
	gorm.Model
	Status         ParticipationStatus `json:"status" gorm:"type:varchar(20);default:'DRAFT';not null"`
	ChallengePoint uint                `json:"challenge_point"` // point in this challenge

	UserID uint           `json:"user_id" gorm:"uniqueIndex:idx_user_challenge"`
	User   userModel.User `json:"user" gorm:"foreignkey:UserID"`

	ChallengeID uint                 `json:"challenge_id" gorm:"uniqueIndex:idx_user_challenge"`
	Challenge   arenaModel.Challenge `json:"challenge" gorm:"foreignkey:ChallengeID"`
	DeletedAt   gorm.DeletedAt       `gorm:"uniqueIndex:idx_user_challenge"`
}

type UserSeasonStat struct {
	gorm.Model
	TotalPoint uint                `json:"total_point"`
	Status     ParticipationStatus `json:"status" gorm:"type:varchar(20);default:'DRAFT';not null"`

	UserID uint           `json:"user_id" gorm:"uniqueIndex:idx_user_season"`
	User   userModel.User `json:"user" gorm:"foreignkey:UserID"`

	SeasonID  uint              `json:"season_id" gorm:"uniqueIndex:idx_user_season"`
	Season    arenaModel.Season `json:"season" gorm:"foreignkey:SeasonID"`
	DeletedAt gorm.DeletedAt    `gorm:"uniqueIndex:idx_user_season"`
}
