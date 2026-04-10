package model

func (round *Round) ToResponse() *RoundResponse {
	return &RoundResponse{
		ID:         round.ID,
		Name:       round.Name,
		ImageUrl:   round.ImageUrl,
		OccurredAt: round.OccurredAt.Format("2006-01-02"), // yyyy-mm-dd
		TotalPoint: round.TotalPoint,
		SeasonID:   round.SeasonID,
		SeasonName: round.Season.Name,
	}
}

func (request *RoundRequest) ToEntity() *Round {
	return &Round{
		Name:       request.Name,
		ImageUrl:   request.ImageUrl,
		OccurredAt: request.OccurredAt,
		TotalPoint: request.TotalPoint,
		SeasonID:   request.SeasonID,
		Status:     request.Status,
	}
}

func (round *Round) UpdateFromRequest(request *RoundRequest) {
	round.Name = request.Name
	round.ImageUrl = request.ImageUrl
	round.OccurredAt = request.OccurredAt
	round.TotalPoint = request.TotalPoint
	round.SeasonID = request.SeasonID
	round.Status = request.Status
}
