package model

type SeasonResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ImageUrl  string `json:"image_url"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}

type RoundResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	ImageUrl   string `json:"image_url"`
	OccurredAt string `json:"occurred_at"`
	TotalPoint uint   `json:"total_point"`
	SeasonID   uint   `json:"season_id"`
	SeasonName string `json:"season_name"`
}

type ChallengeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
	Point       uint   `json:"point"`
	RoundID     uint   `json:"round_id"`
	RoundName   string `json:"round_name"`
	SeasonID    uint   `json:"season_id"`
	SeasonName  string `json:"season_name"`
}
