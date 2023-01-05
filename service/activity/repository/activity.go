package repository

import (
	"absensi-karyawan-api/service/activity/entity"
	"gorm.io/gorm"
)

type ActivityRepository interface {
	CreateActivityRepository(activity entity.Activity) (err error)
	UpdateActivityRepository(activity entity.Activity) (err error)
	GetActivityByIDRepositry(ID int64) (response entity.Activity, err error)
	GetAllActivityRepository() (response []entity.Activity, err error)
	DeleteActivityRepository(activity entity.Activity) (err error)
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *activityRepository {
	return &activityRepository{db}
}

func (r *activityRepository) CreateActivityRepository(activity entity.Activity) (err error) {
	err = r.db.Create(&activity).Error

	return err
}

func (r *activityRepository) UpdateActivityRepository(activity entity.Activity) (err error) {
	err = r.db.Save(&activity).Error

	return err
}

func (r *activityRepository) GetActivityByIDRepositry(ID int64) (response entity.Activity, err error) {
	err = r.db.Where("id = ?", ID).Find(&response).Error

	return response, err
}

func (r *activityRepository) GetAllActivityRepository() (response []entity.Activity, err error) {
	err = r.db.Find(&response).Error

	return response, err
}

func (r *activityRepository) DeleteActivityRepository(activity entity.Activity) (err error) {
	err = r.db.Delete(&activity).Error
	return err
}
