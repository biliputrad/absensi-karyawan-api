package response

import (
	"absensi-karyawan-api/service/activity/response"
	"absensi-karyawan-api/service/attendance/entity"
)

type Attendance struct {
	ID       int64  `json:"id"`
	CheckIn  string `json:"check_in"`
	CheckOut string `json:"check_out"`
	Activity []response.Activity
}

func ConvertEntityToResponseAttendance(attendance entity.Attendance) (result Attendance) {
	result.ID = attendance.ID
	result.CheckIn = attendance.CheckIn.Format("02/01/2006 15:04:05")
	result.CheckOut = attendance.CheckOut.Format("02/01/2006 15:04:05")

	for _, activity := range attendance.Activity {
		convertedActivity := response.ConvertEntityToResponseActivity(activity)
		result.Activity = append(result.Activity, convertedActivity)
	}

	return result
}
