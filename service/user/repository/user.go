package repository

import (
	"absensi-karyawan-api/service/user/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserRepository(user entity.User) (err error)
	GetUserByUsernameRepository(username string) (response entity.User, err error)
	UpdateUserDivisionAndRoleByUserIDRepository(updateUser entity.User) (err error)
	GetUserByIDRepositry(ID int64) (response entity.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUserRepository(user entity.User) (err error) {
	err = r.db.Create(&user).Error

	return err
}

func (r *userRepository) GetUserByUsernameRepository(username string) (response entity.User, err error) {
	err = r.db.Find(&response, "user_name = ?", username).Error

	return response, err
}

func (r *userRepository) UpdateUserDivisionAndRoleByUserIDRepository(updateUser entity.User) (err error) {
	err = r.db.Model(&entity.User{}).Select("role_id", "division_id", "updated_by").Where("id = ?", updateUser.ID).Updates(updateUser).Error

	return err
}

func (r *userRepository) GetUserByIDRepositry(ID int64) (response entity.User, err error) {
	err = r.db.Where("id = ?", ID).Find(&response).Error

	return response, err
}
