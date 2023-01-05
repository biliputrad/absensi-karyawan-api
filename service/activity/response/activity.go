package response

import "absensi-karyawan-api/service/activity/entity"

type Activity struct {
	ID           int64  `json:"id"`
	NameActivity string `json:"name_activity"`
}

func ConvertEntityToResponseActivity(attendance entity.Activity) (response Activity) {
	response.ID = attendance.ID
	response.NameActivity = attendance.NameActivity

	return response
}
