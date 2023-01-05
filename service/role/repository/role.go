package repository

import (
	"absensi-karyawan-api/service/role/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRoleRepository(role entity.Role) (err error)
	GetRoleByRoleNameRepository(roleName string) (response entity.Role, err error)
	GetRoleByIDRepositry(ID int64) (response entity.Role, err error)
	GetAllRoleRepository() (response []entity.Role, err error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) CreateRoleRepository(role entity.Role) (err error) {
	err = r.db.Create(&role).Error

	return err
}

func (r *roleRepository) GetRoleByRoleNameRepository(roleName string) (response entity.Role, err error) {
	err = r.db.Find(&response, "name = ?", roleName).Error

	return response, err
}

func (r *roleRepository) GetRoleByIDRepositry(ID int64) (response entity.Role, err error) {
	err = r.db.Where("id = ?", ID).Find(&response).Error

	return response, err
}

func (r *roleRepository) GetAllRoleRepository() (response []entity.Role, err error) {
	err = r.db.Find(&response).Error

	return response, err
}
