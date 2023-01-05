package usecase

import (
	"absensi-karyawan-api/constant"
	"absensi-karyawan-api/service/attendance/form"
	"absensi-karyawan-api/service/attendance/repository"
	"absensi-karyawan-api/service/attendance/response"
	userRepository "absensi-karyawan-api/service/user/repository"
	"errors"
	"time"
)

type AttendanceUseCase interface {
	CheckInAttendanceUseCase(idUser int64) (result bool, err error)
	CheckOutAttendanceUseCase(idUser int64) (result bool, err error)
	GetAttendanceByIDUseCase(ID int64) (result response.Attendance, err error)
	GetAllAttendanceUseCase(idUser int64) (result []response.Attendance, err error)
	GetAttendanceByDateUseCase(date form.AttedanceDate, idUser int64) (result response.Attendance, err error)
}

type attendanceUseCase struct {
	attendanceRepository repository.AttendanceRepository
	userRepository       userRepository.UserRepository
}

func NewAttendanceUseCase(attendanceRepository repository.AttendanceRepository, userRepository userRepository.UserRepository) *attendanceUseCase {
	return &attendanceUseCase{attendanceRepository, userRepository}
}

func (s *attendanceUseCase) CheckInAttendanceUseCase(idUser int64) (result bool, err error) {
	lastAttendance, err := s.attendanceRepository.GetLastestCheckIn(idUser)

	if err.Error() == constant.RecordNotFound {
		err = nil
	}

	if err != nil {
		return false, err
	}

	if lastAttendance.ID == 0 {
		return false, errors.New(constant.CheckInErr)
	}

	user, err := s.userRepository.GetUserByIDRepositry(idUser)
	if err != nil {
		return false, err
	}

	input := form.ConvertIntoEntityAttendanceCheckIn(user.Email, user.ID)

	err = s.attendanceRepository.CreateAttendanceRepository(input)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *attendanceUseCase) CheckOutAttendanceUseCase(idUser int64) (result bool, err error) {
	user, err := s.userRepository.GetUserByIDRepositry(idUser)
	if err != nil {
		return false, err
	}

	recentAttendance, err := s.attendanceRepository.GetLastestCheckIn(idUser)
	if err != nil {
		return false, err
	}

	if recentAttendance.ID == 0 {
		return false, errors.New(constant.CheckOutErr)
	}

	recentAttendance.CheckOut = time.Now().Local()
	recentAttendance.UpdatedBy = user.Email

	err = s.attendanceRepository.UpdateAttendanceRepository(recentAttendance)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *attendanceUseCase) GetAttendanceByIDUseCase(ID int64) (result response.Attendance, err error) {
	attendance, err := s.attendanceRepository.GetAttendanceByIDRepositry(ID)
	if err != nil {
		return result, err
	}

	result = response.ConvertEntityToResponseAttendance(attendance)

	return result, nil
}

func (s *attendanceUseCase) GetAllAttendanceUseCase(idUser int64) (result []response.Attendance, err error) {
	attendances, err := s.attendanceRepository.GetAllAttendanceRepository(idUser)
	for _, attendance := range attendances {
		responseAttendance := response.ConvertEntityToResponseAttendance(attendance)
		result = append(result, responseAttendance)
	}
	if err != nil {
		return result, err
	}

	return result, err

}

func (s *attendanceUseCase) GetAttendanceByDateUseCase(date form.AttedanceDate, idUser int64) (result response.Attendance, err error) {
	dateParse, err := time.ParseInLocation("02/01/2006", date.Date, time.Local)
	if err != nil {
		return result, err
	}

	attedances, err := s.attendanceRepository.GetAllAttendanceByDateRepository(dateParse, idUser)
	if err != nil {
		return result, err
	}

	result = response.ConvertEntityToResponseAttendance(attedances)

	return result, err
}
