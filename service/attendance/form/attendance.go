package form

import (
	"absensi-karyawan-api/service/attendance/entity"
	"time"
)

type AttedanceDate struct {
	Date string `json:"date" binding:"required"`
}

func ConvertIntoEntityAttendanceCheckIn(email string, userID int64) (response entity.Attendance) {
	response.CheckIn = time.Now().Local()
	response.UserID = userID
	response.CreatedBy = email

	return response
}
