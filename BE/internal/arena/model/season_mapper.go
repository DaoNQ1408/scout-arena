package model

func (season *Season) ToResponse() *SeasonResponse {
	return &SeasonResponse{
		ID:        season.ID,
		Name:      season.Name,
		ImageUrl:  season.ImageUrl,
		StartedAt: season.StartedAt.Format("2006-01-02"), // yyyy-mm-dd
		EndedAt:   season.EndedAt.Format("2006-01-02"),
	}
}

func (request *SeasonRequest) ToEntity() *Season {
	return &Season{
		Name:      request.Name,
		ImageUrl:  request.ImageUrl,
		StartedAt: request.StartedAt,
		EndedAt:   request.EndedAt,
		Status:    request.Status,
	}
}

func (season *Season) UpdateFromRequest(request *SeasonRequest) {
	season.Name = request.Name
	season.ImageUrl = request.ImageUrl
	season.StartedAt = request.StartedAt
	season.EndedAt = request.EndedAt
	season.Status = request.Status
}
