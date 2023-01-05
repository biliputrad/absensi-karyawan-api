package activity

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/activity/repository"
	"absensi-karyawan-api/service/activity/usecase"
	userRepository "absensi-karyawan-api/service/user/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteActivity(db *gorm.DB, routerGroup *gin.RouterGroup) {
	activityRepository := repository.NewActivityRepository(db)
	userRepo := userRepository.NewUserRepository(db)
	activityUseCase := usecase.NewActivityUseCase(activityRepository, userRepo)
	handlerActivity := NewActivityHandler(activityUseCase)

	routerGroup.POST("/create", helper.Auth(), handlerActivity.CreateActivityHandler)
	routerGroup.GET("/get/all", helper.Auth(), handlerActivity.GetAllActivityHandler)
	routerGroup.GET("/get/:id", helper.Auth(), handlerActivity.GetActivityByIDHandler)
	routerGroup.PUT("update", helper.Auth(), handlerActivity.UpdateActivityHandler)
	routerGroup.DELETE("delete/:id", helper.Auth(), handlerActivity.DeleteActivityHandler)
}
