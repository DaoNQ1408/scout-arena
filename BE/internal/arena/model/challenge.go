package model

import (
	"scout-arena/internal/user/model"

	"gorm.io/gorm"
)

type Challenge struct {
	gorm.Model
	Name        string      `json:"name" gorm:"size:50;not null"`
	Description string      `json:"description" gorm:"size:255"`
	ImageUrl    string      `json:"image_url" gorm:"size:255"`
	Point       uint        `json:"point" gorm:"not null"`
	Status      ArenaStatus `json:"status" gorm:"type:varchar(20);default:'DRAFT';not null"`

	RoundID uint  `json:"round_id" gorm:"not null"`
	Round   Round `json:"round" gorm:"foreignKey:RoundID"`

	RankID uint       `json:"rank_id" gorm:"not null, default:1"`
	Rank   model.Rank `json:"rank" gorm:"foreignKey:RankID"`
}

type ChallengeRequest struct { // use for create and update
	Name        string      `json:"name" binding:"required"`
	ImageUrl    string      `json:"image_url" binding:"required,url"`
	Description string      `json:"description" gorm:"size:255"`
	Point       uint        `json:"point" binding:"required,gt=0"`
	RoundID     uint        `json:"round_id" binding:"required"`
	Status      ArenaStatus `json:"status" binding:"required,arena_status"`
	RankID      uint        `json:"rank_id" binding:"required"`
}

type ChallengeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
	Point       uint   `json:"point"`
	RoundID     uint   `json:"round_id"`
	RoundName   string `json:"round_name"`
	SeasonID    uint   `json:"season_id"`
	SeasonName  string `json:"season_name"`
	RankID      uint   `json:"rank_id"`
	RankName    string `json:"rank_name"`
}
