package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"size:25;not null"`
}

type RoleRequest struct {
	Name string `json:"name" binding:"required"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
