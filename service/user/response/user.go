package response

import (
	"absensi-karyawan-api/service/user/entity"
	"time"
)

type LoginResponse struct {
	Type      string    `json:"type"`
	Token     string    `json:"access_token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Division string `json:"division"`
	Role     string `json:"role"`
}

func ConvertEntityToResponseUser(user entity.User, roleName, divisionName string) (response UserResponse) {
	response.ID = user.ID
	response.Email = user.Email
	response.UserName = user.UserName
	response.Name = user.Name
	response.Division = roleName
	response.Role = divisionName

	return response
}
