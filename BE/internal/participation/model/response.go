package model

type UserChallengeResponse struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
	ChallengeID uint   `json:"challenge_id"`
	Challenge   string `json:"challenge"`

	Status         ParticipationStatus `json:"status"`
	ChallengePoint uint                `json:"challenge_point"`
}

type UserRoundResponse struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	RoundID  uint   `json:"round_id"`
	Round    string `json:"round"`

	Status     ParticipationStatus `json:"status"`
	RoundPoint uint                `json:"round_point"`
	TotalPoint uint                `json:"total_point"`
}

type UserSeasonResponse struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	SeasonID uint   `json:"season_id"`
	Season   string `json:"season"`

	Status     ParticipationStatus `json:"status"`
	TotalPoint uint                `json:"total_point"`
}
