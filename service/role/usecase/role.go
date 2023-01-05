package usecase

import (
	"absensi-karyawan-api/service/role/form"
	"absensi-karyawan-api/service/role/repository"
	"absensi-karyawan-api/service/role/response"
	userRepository "absensi-karyawan-api/service/user/repository"
)

type RoleUseCase interface {
	CreateRoleUseCase(input form.Role, idUser int64) (result bool, err error)
	GetRoleByIDUseCase(ID int64) (result response.Role, err error)
	GetAllRoleUseCase() (result []response.Role, err error)
}

type roleUseCase struct {
	roleRepository repository.RoleRepository
	userRepository userRepository.UserRepository
}

func NewRoleUseCase(roleRepository repository.RoleRepository, userRepository userRepository.UserRepository) *roleUseCase {
	return &roleUseCase{roleRepository, userRepository}
}

func (s *roleUseCase) CreateRoleUseCase(input form.Role, idUser int64) (result bool, err error) {
	user, err := s.userRepository.GetUserByIDRepositry(idUser)
	if err != nil {
		return false, err
	}

	role := form.ConvertIntoEntityRole(input, user.Email)

	err = s.roleRepository.CreateRoleRepository(role)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *roleUseCase) GetRoleByIDUseCase(ID int64) (result response.Role, err error) {
	role, err := s.roleRepository.GetRoleByIDRepositry(ID)
	if err != nil {
		return result, err
	}

	result = response.ConvertEntityToResponseRole(role)

	return result, nil
}

func (s *roleUseCase) GetAllRoleUseCase() (result []response.Role, err error) {
	roles, err := s.roleRepository.GetAllRoleRepository()
	for _, role := range roles {
		responseRole := response.ConvertEntityToResponseRole(role)
		result = append(result, responseRole)
	}
	if err != nil {
		return result, err
	}

	return result, err

}
