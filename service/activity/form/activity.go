package form

import "absensi-karyawan-api/service/activity/entity"

type Activity struct {
	NameActivity string `json:"name_activity" binding:"required"`
	AttendanceID int64  `json:"attendance_id" binding:"required"`
}

type UpdateActivity struct {
	ActivityID   int64  `json:"activity_id" binding:"required"`
	NameActivity string `json:"name_activity" binding:"required"`
}

func ConvertIntoEntityActivity(activity Activity, email string) (response entity.Activity) {
	response.NameActivity = activity.NameActivity
	response.CreatedBy = email
	response.AttendanceID = activity.AttendanceID

	return response
}
