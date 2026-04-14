package model

func (rank *Rank) ToResponse() *RankResponse {
	return &RankResponse{
		ID:   rank.ID,
		Name: rank.Name,
	}
}

func (request *RankRequest) ToEntity() *Rank {
	return &Rank{
		Name: request.Name,
	}
}

func (rank *Rank) UpdateFromRequest(request *RankRequest) {
	rank.Name = request.Name
}
