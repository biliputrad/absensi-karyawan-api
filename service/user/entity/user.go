package entity

import (
	"absensi-karyawan-api/helper"
	entityAttendance "absensi-karyawan-api/service/attendance/entity"
)

type User struct {
	helper.Base
	Email      string `gorm:"unique"`
	UserName   string `gorm:"unique"`
	Password   string
	Name       string
	RoleID     int64
	DivisionID int64
	Attendance []entityAttendance.Attendance `gorm:"foreignKey:UserID"`
}
