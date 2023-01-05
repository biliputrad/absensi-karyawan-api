package activity

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/attendance/repository"
	"absensi-karyawan-api/service/attendance/usecase"
	userRepository "absensi-karyawan-api/service/user/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteAttendance(db *gorm.DB, routerGroup *gin.RouterGroup) {
	attendanceRepository := repository.NewAttendanceRepository(db)
	userRepo := userRepository.NewUserRepository(db)
	attendanceUseCase := usecase.NewAttendanceUseCase(attendanceRepository, userRepo)
	handlerAttendance := NewAttendanceHandler(attendanceUseCase)

	routerGroup.POST("/check-in", helper.Auth(), handlerAttendance.CreateAttedanceCheckInHandler)
	routerGroup.PUT("/check-out", helper.Auth(), handlerAttendance.CreateAttedanceCheckOutHandler)
	routerGroup.GET("/get/all", helper.Auth(), handlerAttendance.GetAllAttendanceHandler)
	routerGroup.GET("/get/:id", helper.Auth(), handlerAttendance.GetAttendanceByIDHandler)
	routerGroup.GET("get/date", helper.Auth(), handlerAttendance.GetAttendanceByDateHandler)
}
