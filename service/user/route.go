package user

import (
	"absensi-karyawan-api/helper"
	repositoryDivision "absensi-karyawan-api/service/division/repository"
	repositoryRole "absensi-karyawan-api/service/role/repository"
	"absensi-karyawan-api/service/user/repository"
	"absensi-karyawan-api/service/user/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteUser(db *gorm.DB, routerGroup *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	roleRepository := repositoryRole.NewRoleRepository(db)
	divisionRepository := repositoryDivision.NewDivisionRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository, divisionRepository, roleRepository)
	handlerUser := NewUserHandler(userUseCase)

	routerGroup.POST("/create", handlerUser.CreateUserHandler)
	routerGroup.POST("/login", handlerUser.LoginUserHandler)
	routerGroup.PUT("/update", helper.Auth(), handlerUser.UpdateUserRoleAndDivisionHandler)
	routerGroup.GET("/get", helper.Auth(), handlerUser.GetUserByIDHandler)
	routerGroup.POST("/logout", helper.Auth(), handlerUser.LogoutUserHandler)
}
