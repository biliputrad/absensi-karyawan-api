package entity

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/activity/entity"
	"time"
)

type Attendance struct {
	helper.Base
	CheckIn  time.Time
	CheckOut time.Time
	UserID   int64
	Activity []entity.Activity `gorm:"foreignKey:AttendanceID"`
}
