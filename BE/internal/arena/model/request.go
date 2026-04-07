package model

import "time"

type SeasonRequest struct { // use for create and update
	Name      string      `json:"name" binding:"required"`
	ImageUrl  string      `json:"image_url" binding:"required,url"`
	StartedAt time.Time   `json:"started_at" binding:"required"`
	EndedAt   time.Time   `json:"ended_at" binding:"required,gtfield=StartedAt"`
	Status    ArenaStatus `json:"status" binding:"required,arena_status"`
}

type RoundRequest struct { // use for create and update
	Name       string      `json:"name" binding:"required"`
	ImageUrl   string      `json:"image_url" binding:"required,url"`
	OccurredAt time.Time   `json:"occurred_at" binding:"required"`
	TotalPoint uint        `json:"total_point" binding:"required,gt=0"`
	SeasonID   uint        `json:"season_id" binding:"required"`
	Status     ArenaStatus `json:"status" binding:"required,arena_status"`
}

type ChallengeRequest struct { // use for create and update
	Name        string      `json:"name" binding:"required"`
	ImageUrl    string      `json:"image_url" binding:"required,url"`
	Description string      `json:"description" gorm:"size:255"`
	Point       uint        `json:"point" binding:"required,gt=0"`
	RoundID     uint        `json:"round_id" binding:"required"`
	Status      ArenaStatus `json:"status" binding:"required,arena_status"`
}
