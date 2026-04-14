package model

import "gorm.io/gorm"

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

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`

	RankID uint `json:"rank_id" binding:"required"`
	TeamID uint `json:"team_id" binding:"required"`
	RoleID uint `json:"role_id" binding:"required"`
}

type UpdateUserRequest struct {
	Name   string     `json:"name" binding:"required"`
	Status UserStatus `json:"status" binding:"required,user_status"`

	RankID uint `json:"rank_id" binding:"required"`
	TeamID uint `json:"team_id" binding:"required"`
	RoleID uint `json:"role_id" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	Username    string `json:"username" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type LoginResponse struct {
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`

	RankID uint   `json:"rank_id"`
	Rank   string `json:"rank"`

	TeamID uint   `json:"team_id"`
	Team   string `json:"team"`

	RoleID uint   `json:"role_id"`
	Role   string `json:"role"`
}
