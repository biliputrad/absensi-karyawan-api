package division

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/division/repository"
	"absensi-karyawan-api/service/division/usecase"
	userRepository "absensi-karyawan-api/service/user/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteDivision(db *gorm.DB, routerGroup *gin.RouterGroup) {
	divisionRepository := repository.NewDivisionRepository(db)
	userRepo := userRepository.NewUserRepository(db)
	divisionUseCase := usecase.NewDivisionUseCase(divisionRepository, userRepo)
	handlerDivision := NewDivisionHandler(divisionUseCase)

	routerGroup.POST("/create", helper.Auth(), handlerDivision.CreateDivisionHandler)
	routerGroup.GET("/get/all", helper.Auth(), handlerDivision.GetAllDivisionHandler)
}
