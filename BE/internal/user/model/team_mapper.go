package model

func (team *Team) ToResponse() *TeamResponse {
	return &TeamResponse{
		ID:     team.ID,
		Name:   team.Name,
		Status: string(team.Status),
	}
}

func (request *TeamRequest) ToEntity() *Team {
	return &Team{
		Name:   request.Name,
		Status: request.Status,
	}
}

func (team *Team) UpdateFromRequest(request *TeamRequest) {
	team.Name = request.Name
	team.Status = request.Status
}
