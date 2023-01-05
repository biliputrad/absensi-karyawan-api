package usecase

import (
	"absensi-karyawan-api/service/activity/form"
	"absensi-karyawan-api/service/activity/repository"
	"absensi-karyawan-api/service/activity/response"
	userRepository "absensi-karyawan-api/service/user/repository"
)

type ActivityUseCase interface {
	CreateActivityUseCase(activity form.Activity, idUser int64) (result bool, err error)
	UpdateActivityUseCase(activity form.UpdateActivity, idUser int64) (result bool, err error)
	GetActivityByIDUseCase(ID int64) (result response.Activity, err error)
	GetAllActivityUseCase() (result []response.Activity, err error)
	DeleteActivityUseCase(ID int64) (result bool, err error)
}

type activityUseCase struct {
	activityRepository repository.ActivityRepository
	userRepository     userRepository.UserRepository
}

func NewActivityUseCase(activityRepository repository.ActivityRepository, userRepository userRepository.UserRepository) *activityUseCase {
	return &activityUseCase{activityRepository, userRepository}
}

func (s *activityUseCase) CreateActivityUseCase(input form.Activity, idUser int64) (result bool, err error) {
	user, err := s.userRepository.GetUserByIDRepositry(idUser)
	if err != nil {
		return false, err
	}

	activity := form.ConvertIntoEntityActivity(input, user.Email)

	err = s.activityRepository.CreateActivityRepository(activity)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *activityUseCase) UpdateActivityUseCase(activity form.UpdateActivity, idUser int64) (result bool, err error) {
	user, err := s.userRepository.GetUserByIDRepositry(idUser)
	if err != nil {
		return false, err
	}

	recentActivity, err := s.activityRepository.GetActivityByIDRepositry(activity.ActivityID)
	if err != nil {
		return false, err
	}

	recentActivity.UpdatedBy = user.Email
	recentActivity.NameActivity = activity.NameActivity

	err = s.activityRepository.UpdateActivityRepository(recentActivity)
	if err != nil {
		return false, err
	}

	return true, err
}

func (s *activityUseCase) GetActivityByIDUseCase(ID int64) (result response.Activity, err error) {
	activity, err := s.activityRepository.GetActivityByIDRepositry(ID)
	if err != nil {
		return result, err
	}

	result = response.ConvertEntityToResponseActivity(activity)

	return result, nil
}

func (s *activityUseCase) GetAllActivityUseCase() (result []response.Activity, err error) {
	activitys, err := s.activityRepository.GetAllActivityRepository()
	for _, activity := range activitys {
		responseActivity := response.ConvertEntityToResponseActivity(activity)
		result = append(result, responseActivity)
	}
	if err != nil {
		return result, err
	}

	return result, err

}

func (s *activityUseCase) DeleteActivityUseCase(ID int64) (result bool, err error) {
	activity, err := s.activityRepository.GetActivityByIDRepositry(ID)
	if err != nil {
		return false, err
	}

	err = s.activityRepository.DeleteActivityRepository(activity)
	if err != nil {
		return false, err
	}

	return true, err
}
