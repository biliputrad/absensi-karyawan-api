package response

import "absensi-karyawan-api/service/role/entity"

type Role struct {
	ID   int64  `json:"role_id"`
	Name string `json:"name"`
}

func ConvertEntityToResponseRole(role entity.Role) (response Role) {
	response.ID = role.ID
	response.Name = role.Name

	return response
}
