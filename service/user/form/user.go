package form

import "absensi-karyawan-api/service/user/entity"

type CreateUser struct {
	Email      string `json:"email" binding:"required"`
	UserName   string `json:"user_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Name       string `json:"name" binding:"required"`
	DivisionID int64  `json:"division_id"`
	RoleID     int64  `json:"role_id"`
}

type LoginUser struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRoleAndDivision struct {
	ID         int64
	Email      string
	DivisionID int64 `json:"division_id" binding:"required"`
	RoleID     int64 `json:"role_id" binding:"required"`
}

func ConvertIntoEntityUser(user CreateUser) (response entity.User) {
	response.Email = user.Email
	response.UserName = user.UserName
	response.Name = user.Name

	return response
}
