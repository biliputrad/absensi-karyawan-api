package response

import "absensi-karyawan-api/service/division/entity"

type Division struct {
	ID   int64  `json:"division_id"`
	Name string `json:"division_name"`
}

func ConvertEntityToResponseDivision(division entity.Division) (response Division) {
	response.ID = division.ID
	response.Name = division.Name

	return response
}
