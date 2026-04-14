package model

import "gorm.io/gorm"

type Rank struct {
	gorm.Model
	Name string `json:"name" gorm:"size:25;not null"`
}

type RankRequest struct {
	Name string `json:"name" binding:"required"`
}

type RankResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
