package model

import (
	"gorm.io/gorm"
)

type Rank struct {
	gorm.Model
	Name string `json:"name" gorm:"size:25;not null"`
}

type Team struct {
	gorm.Model
	Name   string     `json:"name" gorm:"size:25;not null"`
	Status TeamStatus `json:"status" gorm:"type:varchar(20);default:'ACTIVE';not null"`
}

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"size:25;not null"`
}

type User struct {
	gorm.Model
	Username string     `json:"username" gorm:"uniqueIndex;not null"`
	Password string     `json:"password" gorm:"not null"`
	Name     string     `json:"name" gorm:"size:50;not null"`
	Status   UserStatus `json:"status" gorm:"type:varchar(20);default:'ACTIVE';not null"`

	RankID uint `json:"rank_id"`
	Rank   Rank `json:"rank" gorm:"foreignKey:RankID"`

	TeamID uint `json:"team_id"`
	Team   Team `json:"team" gorm:"foreignKey:TeamID"`

	RoleID uint `json:"role_id"`
	Role   Role `json:"role" gorm:"foreignKey:RoleID"`
}
