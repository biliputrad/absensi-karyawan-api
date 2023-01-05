package form

import "absensi-karyawan-api/service/role/entity"

type Role struct {
	RoleName string `json:"role_name" binding:"required"`
}

func ConvertIntoEntityRole(role Role, email string) (response entity.Role) {
	response.Name = role.RoleName
	response.CreatedBy = email

	return response
}
