package model

func (challenge *Challenge) ToResponse() *ChallengeResponse {
	return &ChallengeResponse{
		ID:          challenge.ID,
		Name:        challenge.Name,
		ImageUrl:    challenge.ImageUrl,
		Description: challenge.Description,
		Point:       challenge.Point,
		RoundID:     challenge.RoundID,
		RoundName:   challenge.Round.Name,
		SeasonID:    challenge.Round.SeasonID,
		SeasonName:  challenge.Round.Season.Name,
	}
}

func (request *ChallengeRequest) ToEntity() *Challenge {
	return &Challenge{
		Name:        request.Name,
		ImageUrl:    request.ImageUrl,
		Description: request.Description,
		Point:       request.Point,
		RoundID:     request.RoundID,
		Status:      request.Status,
	}
}

func (challenge *Challenge) UpdateFromRequest(request *ChallengeRequest) {
	challenge.Name = request.Name
	challenge.Description = request.Description
	challenge.Point = request.Point
	challenge.ImageUrl = request.ImageUrl
	challenge.RoundID = request.RoundID
	challenge.Status = request.Status
}
