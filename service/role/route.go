package role

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/role/repository"
	"absensi-karyawan-api/service/role/usecase"
	userRepository "absensi-karyawan-api/service/user/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteRole(db *gorm.DB, routerGroup *gin.RouterGroup) {
	roleRepository := repository.NewRoleRepository(db)
	userRepo := userRepository.NewUserRepository(db)
	roleUseCase := usecase.NewRoleUseCase(roleRepository, userRepo)
	handlerRole := NewRoleHandler(roleUseCase)

	routerGroup.POST("/create", helper.Auth(), handlerRole.CreateRoleHandler)
	routerGroup.GET("/get/all", helper.Auth(), handlerRole.GetAllRoleHandler)
}
