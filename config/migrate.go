package config

import (
	entityActivity "absensi-karyawan-api/service/activity/entity"
	entityAttendance "absensi-karyawan-api/service/attendance/entity"
	entityDivision "absensi-karyawan-api/service/division/entity"
	entityRole "absensi-karyawan-api/service/role/entity"
	entityUser "absensi-karyawan-api/service/user/entity"
	"gorm.io/gorm"
)

type Entity struct {
	EntityActivity   interface{}
	EntityAttendance interface{}
	EntityUser       interface{}
	EntityRole       interface{}
	EntityDivision   interface{}
}

func RegisterEntityActivity() []Entity {
	return []Entity{
		{EntityActivity: entityActivity.Activity{}},
	}
}

func RegisterEntityAttendance() []Entity {
	return []Entity{
		{EntityAttendance: entityAttendance.Attendance{}},
	}
}

func RegisterEntityUser() []Entity {
	return []Entity{
		{EntityUser: entityUser.User{}},
	}
}

func RegisterEntityRole() []Entity {
	return []Entity{
		{EntityRole: entityRole.Role{}},
	}
}

func RegisterEntityDivision() []Entity {
	return []Entity{
		{EntityDivision: entityDivision.Division{}},
	}
}

func Activity(db *gorm.DB) error {
	for _, entity := range RegisterEntityActivity() {
		dbMigErr := db.Debug().AutoMigrate(entity.EntityActivity)
		if dbMigErr != nil {
			return dbMigErr
		}
	}
	return nil
}

func Attendance(db *gorm.DB) error {
	for _, entity := range RegisterEntityAttendance() {
		dbMigErr := db.Debug().AutoMigrate(entity.EntityAttendance)
		if dbMigErr != nil {
			return dbMigErr
		}
	}
	return nil
}

func User(db *gorm.DB) error {
	for _, entity := range RegisterEntityUser() {
		dbMigErr := db.Debug().AutoMigrate(entity.EntityUser)
		if dbMigErr != nil {
			return dbMigErr
		}
	}
	return nil
}

func Role(db *gorm.DB) error {
	for _, entity := range RegisterEntityRole() {
		dbMigErr := db.Debug().AutoMigrate(entity.EntityRole)
		if dbMigErr != nil {
			return dbMigErr
		}
	}
	return nil
}

func Division(db *gorm.DB) error {
	for _, entity := range RegisterEntityDivision() {
		dbMigErr := db.Debug().AutoMigrate(entity.EntityDivision)
		if dbMigErr != nil {
			return dbMigErr
		}
	}
	return nil
}
