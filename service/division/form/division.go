package form

import "absensi-karyawan-api/service/division/entity"

type Division struct {
	Name string `json:"division_name" binding:"required"`
}

func ConvertIntoEntityDivision(division Division, email string) (response entity.Division) {
	response.Name = division.Name
	response.CreatedBy = email

	return response
}
