package model

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
