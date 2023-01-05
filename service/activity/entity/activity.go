package entity

import (
	"absensi-karyawan-api/helper"
)

type Activity struct {
	helper.Base
	NameActivity string
	AttendanceID int64
}
