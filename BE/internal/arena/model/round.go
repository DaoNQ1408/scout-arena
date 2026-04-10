package model

import (
	"time"

	"gorm.io/gorm"
)

type Round struct {
	gorm.Model
	Name       string      `json:"name" gorm:"size:100;not null"`
	ImageUrl   string      `json:"image_url" gorm:"size:255"`
	OccurredAt time.Time   `json:"occurred_at" gorm:"type:date;not null"`
	TotalPoint uint        `json:"total_point" gorm:"not null"`
	Status     ArenaStatus `json:"status" gorm:"type:varchar(20);default:'DRAFT';not null"`

	SeasonID uint   `json:"season_id" gorm:"not null"`
	Season   Season `json:"season" gorm:"foreignKey:SeasonID"`
}

type RoundRequest struct { // use for create and update
	Name       string      `json:"name" binding:"required"`
	ImageUrl   string      `json:"image_url" binding:"required,url"`
	OccurredAt time.Time   `json:"occurred_at" binding:"required"`
	TotalPoint uint        `json:"total_point" binding:"required,gt=0"`
	SeasonID   uint        `json:"season_id" binding:"required"`
	Status     ArenaStatus `json:"status" binding:"required,arena_status"`
}

type RoundResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	ImageUrl   string `json:"image_url"`
	OccurredAt string `json:"occurred_at"`
	TotalPoint uint   `json:"total_point"`
	SeasonID   uint   `json:"season_id"`
	SeasonName string `json:"season_name"`
}
