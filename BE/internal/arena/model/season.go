package model

import (
	"time"

	"gorm.io/gorm"
)

type Season struct {
	gorm.Model
	Name      string      `json:"name" gorm:"size:50;not null"`
	ImageUrl  string      `json:"image_url" gorm:"size:255"`
	StartedAt time.Time   `json:"started_at" gorm:"type:date;not null"`
	EndedAt   time.Time   `json:"ended_at" gorm:"type:date;not null"`
	Status    ArenaStatus `json:"status" gorm:"type:varchar(20);default:'DRAFT';not null"`
}

type SeasonRequest struct { // use for create and update
	Name      string      `json:"name" binding:"required"`
	ImageUrl  string      `json:"image_url" binding:"required,url"`
	StartedAt time.Time   `json:"started_at" binding:"required"`
	EndedAt   time.Time   `json:"ended_at" binding:"required,gtfield=StartedAt"`
	Status    ArenaStatus `json:"status" binding:"required,arena_status"`
}

type SeasonResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ImageUrl  string `json:"image_url"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}
