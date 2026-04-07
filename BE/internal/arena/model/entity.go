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

type Challenge struct {
	gorm.Model
	Name        string      `json:"name" gorm:"size:50;not null"`
	Description string      `json:"description" gorm:"size:255"`
	ImageUrl    string      `json:"image_url" gorm:"size:255"`
	Point       uint        `json:"point" gorm:"not null"`
	Status      ArenaStatus `json:"status" gorm:"type:varchar(20);default:'DRAFT';not null"`

	RoundID uint  `json:"round_id" gorm:"not null"`
	Round   Round `json:"round" gorm:"foreignKey:RoundID"`
}
