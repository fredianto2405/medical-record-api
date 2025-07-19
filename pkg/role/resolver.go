package role

import "medical-record-api/internal/constant"

func Resolve(role string) int {
	switch role {
	case constant.RoleOwner:
		return constant.RoleIdOwner
	case constant.RoleManager:
		return constant.RoleIdManager
	case constant.RoleAdmin:
		return constant.RoleIdAdmin
	case constant.RoleDoctor:
		return constant.RoleIdDoctor
	default:
		return constant.InvalidRoleId
	}
}
