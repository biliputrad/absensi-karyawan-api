package repository

import (
	"absensi-karyawan-api/service/division/entity"
	"gorm.io/gorm"
)

type DivisionRepository interface {
	CreateDivisionRepository(division entity.Division) (err error)
	GetDivisionByDivisionNameRepository(divisionName string) (response entity.Division, err error)
	GetDivisionByIDRepositry(ID int64) (response entity.Division, err error)
	GetAllDivisionRepository() (response []entity.Division, err error)
}

type divisionRepository struct {
	db *gorm.DB
}

func NewDivisionRepository(db *gorm.DB) *divisionRepository {
	return &divisionRepository{db}
}

func (r *divisionRepository) CreateDivisionRepository(division entity.Division) (err error) {
	err = r.db.Create(&division).Error

	return err
}

func (r *divisionRepository) GetDivisionByDivisionNameRepository(divisionName string) (response entity.Division, err error) {
	err = r.db.Find(&response, "name = ?", divisionName).Error

	return response, err
}

func (r *divisionRepository) GetDivisionByIDRepositry(ID int64) (response entity.Division, err error) {
	err = r.db.Where("id = ?", ID).Find(&response).Error

	return response, err
}

func (r *divisionRepository) GetAllDivisionRepository() (response []entity.Division, err error) {
	err = r.db.Find(&response).Error

	return response, err
}
