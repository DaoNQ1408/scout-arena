package model

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

type RankRequest struct {
	Name string `json:"name" binding:"required"`
}

type TeamRequest struct {
	Name   string     `json:"name" binding:"required"`
	Status TeamStatus `json:"status" binding:"required,team_status"`
}

type RoleRequest struct {
	Name string `json:"name" binding:"required"`
}
