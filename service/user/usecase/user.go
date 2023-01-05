package usecase

import (
	"absensi-karyawan-api/constant"
	"absensi-karyawan-api/helper"
	divisionRepository "absensi-karyawan-api/service/division/repository"
	roleRepository "absensi-karyawan-api/service/role/repository"
	"absensi-karyawan-api/service/user/entity"
	"absensi-karyawan-api/service/user/form"
	"absensi-karyawan-api/service/user/repository"
	"absensi-karyawan-api/service/user/response"
	"errors"
)

type UserUseCase interface {
	CreateUserUseCase(user form.CreateUser) (result bool, err error)
	LoginUserUseCase(user form.LoginUser) (result response.LoginResponse, err error)
	LogoutUserUseCase() (result response.LoginResponse, err error)
	GetUserByIDUseCase(ID int64) (result response.UserResponse, err error)
	UpdateUserDivisionAndRoleByUserIDUseCase(user form.UpdateUserRoleAndDivision, ID int64) (result bool, err error)
}

type userUseCase struct {
	userRepository     repository.UserRepository
	divisionRepository divisionRepository.DivisionRepository
	roleRepository     roleRepository.RoleRepository
}

func NewUserUseCase(userRepository repository.UserRepository, divisionRepository divisionRepository.DivisionRepository, roleRepository roleRepository.RoleRepository) *userUseCase {
	return &userUseCase{userRepository, divisionRepository, roleRepository}
}

func (s *userUseCase) CreateUserUseCase(input form.CreateUser) (result bool, err error) {
	user := form.ConvertIntoEntityUser(input)
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return false, err
	}

	user.Password = hashedPassword
	user.CreatedBy = user.UserName

	err = s.userRepository.CreateUserRepository(user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *userUseCase) LoginUserUseCase(input form.LoginUser) (result response.LoginResponse, err error) {
	user, err := s.userRepository.GetUserByUsernameRepository(input.UserName)
	if err != nil {
		return result, err
	}

	role, err := s.roleRepository.GetRoleByIDRepositry(user.RoleID)
	if err != nil {
		return result, err
	}

	division, err := s.divisionRepository.GetDivisionByIDRepositry(user.DivisionID)
	if err != nil {
		return result, err
	}

	wrongPassword := helper.ComparePassword(user.Password, input.Password)

	if wrongPassword {
		err = errors.New(constant.WrongPassword)
		return result, err
	}

	token, expired, err := helper.GenerateToken(helper.JwtClaim{
		UserID:   user.ID,
		UserName: user.UserName,
		Division: division.Name,
		Role:     role.Name,
	})

	if err != nil {
		return result, err
	}

	result.Type = constant.Bearer
	result.Token = token
	result.ExpiredAt = expired

	return result, nil

}

func (s *userUseCase) LogoutUserUseCase() (result response.LoginResponse, err error) {
	token, expired, err := helper.GenerateLogoutToken()

	if err != nil {
		return result, err
	}

	result.Type = constant.Bearer
	result.Token = token
	result.ExpiredAt = expired

	return result, nil
}

func (s *userUseCase) GetUserByIDUseCase(ID int64) (result response.UserResponse, err error) {
	user, err := s.userRepository.GetUserByIDRepositry(ID)
	if err != nil {
		return result, err
	}

	role, err := s.roleRepository.GetRoleByIDRepositry(user.RoleID)
	if err != nil {
		return result, err
	}

	divison, err := s.divisionRepository.GetDivisionByIDRepositry(user.DivisionID)
	if err != nil {
		return result, err
	}

	result = response.ConvertEntityToResponseUser(user, role.Name, divison.Name)

	return result, nil
}

func (s *userUseCase) UpdateUserDivisionAndRoleByUserIDUseCase(user form.UpdateUserRoleAndDivision, ID int64) (result bool, err error) {
	dataUser, err := s.GetUserByIDUseCase(ID)
	if err != nil {
		return false, err
	}

	updateUser := entity.User{
		RoleID:     user.RoleID,
		DivisionID: user.DivisionID,
	}
	updateUser.UpdatedBy = dataUser.Email
	updateUser.ID = user.ID

	err = s.userRepository.UpdateUserDivisionAndRoleByUserIDRepository(updateUser)
	if err != nil {
		return false, err
	}

	return true, nil
}
