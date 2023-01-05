package usecase

import (
	"absensi-karyawan-api/service/division/form"
	"absensi-karyawan-api/service/division/repository"
	"absensi-karyawan-api/service/division/response"
	userRepository "absensi-karyawan-api/service/user/repository"
)

type DivisionUseCase interface {
	CreateDivisionUseCase(input form.Division, idUser int64) (result bool, err error)
	GetDivisionByIDUseCase(ID int64) (result response.Division, err error)
	GetAllDivisionUseCase() (result []response.Division, err error)
}

type divisionUseCase struct {
	divisionRepository repository.DivisionRepository
	userRepository     userRepository.UserRepository
}

func NewDivisionUseCase(divisionRepository repository.DivisionRepository, userRepository userRepository.UserRepository) *divisionUseCase {
	return &divisionUseCase{divisionRepository, userRepository}
}

func (s *divisionUseCase) CreateDivisionUseCase(input form.Division, idUser int64) (result bool, err error) {
	user, err := s.userRepository.GetUserByIDRepositry(idUser)
	if err != nil {
		return false, err
	}

	division := form.ConvertIntoEntityDivision(input, user.Email)

	err = s.divisionRepository.CreateDivisionRepository(division)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *divisionUseCase) GetDivisionByIDUseCase(ID int64) (result response.Division, err error) {
	division, err := s.divisionRepository.GetDivisionByIDRepositry(ID)
	if err != nil {
		return result, err
	}

	result = response.ConvertEntityToResponseDivision(division)

	return result, nil
}

func (s *divisionUseCase) GetAllDivisionUseCase() (result []response.Division, err error) {
	divisions, err := s.divisionRepository.GetAllDivisionRepository()
	for _, division := range divisions {
		responseDivision := response.ConvertEntityToResponseDivision(division)
		result = append(result, responseDivision)
	}
	if err != nil {
		return result, err
	}

	return result, err

}
