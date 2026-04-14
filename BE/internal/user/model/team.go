package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name   string     `json:"name" gorm:"size:25;not null"`
	Status TeamStatus `json:"status" gorm:"type:varchar(20);default:'ACTIVE';not null"`
}

type TeamRequest struct {
	Name   string     `json:"name" binding:"required"`
	Status TeamStatus `json:"status" binding:"required,team_status"`
}

type TeamResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
