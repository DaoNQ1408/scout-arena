package model

func (role *Role) ToResponse() *RoleResponse {
	return &RoleResponse{
		ID:   role.ID,
		Name: role.Name,
	}
}

func (request *RoleRequest) ToEntity() *Role {
	return &Role{
		Name: request.Name,
	}
}

func (role *Role) UpdateFromRequest(request *RoleRequest) {
	role.Name = request.Name
}
