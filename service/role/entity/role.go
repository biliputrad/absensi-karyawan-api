package entity

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/user/entity"
)

type Role struct {
	helper.Base
	Name string        `gorm:"unique"`
	User []entity.User `gorm:"foreignKey: RoleID"`
}
