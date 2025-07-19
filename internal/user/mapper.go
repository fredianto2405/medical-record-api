package user

import "medical-record-api/pkg/role"

func MapToEntity(request *Request) *Entity {
	roleID := role.Resolve(request.Role)
	return &Entity{
		Email:     request.Email,
		Password:  request.Password,
		IsActive:  request.IsActive,
		RoleID:    roleID,
		CreatedAt: nil,
		DeletedAt: nil,
	}
}
