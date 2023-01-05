package repository

import (
	"absensi-karyawan-api/service/attendance/entity"
	"gorm.io/gorm"
	"time"
)

type AttendanceRepository interface {
	CreateAttendanceRepository(attendance entity.Attendance) (err error)
	UpdateAttendanceRepository(attendance entity.Attendance) (err error)
	GetAttendanceByIDRepositry(ID int64) (response entity.Attendance, err error)
	GetAllAttendanceRepository(ID int64) (response []entity.Attendance, err error)
	DeleteAttendanceRepository(attendance entity.Attendance) (err error)
	GetAllAttendanceByDateRepository(date time.Time, idUser int64) (response entity.Attendance, err error)
	GetLastestCheckIn(idUser int64) (response entity.Attendance, err error)
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *attendanceRepository {
	return &attendanceRepository{db}
}

func (r *attendanceRepository) CreateAttendanceRepository(attendance entity.Attendance) (err error) {
	err = r.db.Create(&attendance).Error

	return err
}

func (r *attendanceRepository) UpdateAttendanceRepository(attendance entity.Attendance) (err error) {
	err = r.db.Save(&attendance).Error

	return err
}

func (r *attendanceRepository) GetAttendanceByIDRepositry(ID int64) (response entity.Attendance, err error) {
	err = r.db.Preload("Activity").Where("id = ?", ID).Find(&response).Error

	return response, err
}

func (r *attendanceRepository) GetAllAttendanceRepository(ID int64) (response []entity.Attendance, err error) {
	err = r.db.Preload("Activity").Where("user_id = ?", ID).Find(&response).Error

	return response, err
}

func (r *attendanceRepository) DeleteAttendanceRepository(attendance entity.Attendance) (err error) {
	err = r.db.Delete(&attendance).Error
	return err
}

func (r *attendanceRepository) GetAllAttendanceByDateRepository(date time.Time, idUser int64) (response entity.Attendance, err error) {
	err = r.db.Preload("Activity").Where("user_id = ?", idUser).Where("check_in >= ?", date).Find(&response).Error

	return response, err
}

func (r *attendanceRepository) GetLastestCheckIn(idUser int64) (response entity.Attendance, err error) {
	err = r.db.Where("user_id = ?", idUser).Where("check_out = ?", time.Time{}).Last(&response).Error

	return response, err
}
