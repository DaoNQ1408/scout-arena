package model

type CreateUserChallengeRequest struct {
	UserID      uint                `json:"user_id" binding:"required"`
	ChallengeID uint                `json:"challenge_id" binding:"required"`
	Status      ParticipationStatus `json:"status" binding:"required"`
}

type UpdateUserChallengeRequest struct {
	UserChallengeID uint                `json:"user_challenge_id" binding:"required"`
	Status          ParticipationStatus `json:"status" binding:"required"`
}
