package user

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/user/form"
	"absensi-karyawan-api/service/user/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *userHandler {
	return &userHandler{userUseCase}
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var input form.CreateUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := h.userUseCase.CreateUserUseCase(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), result))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var input form.LoginUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := h.userUseCase.LoginUserUseCase(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)

}

func (h *userHandler) LogoutUserHandler(c *gin.Context) {
	result, err := h.userUseCase.LogoutUserUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)

}

func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.GetInt64("id")

	result, err := h.userUseCase.GetUserByIDUseCase(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) UpdateUserRoleAndDivisionHandler(c *gin.Context) {
	var input form.UpdateUserRoleAndDivision
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	id := c.GetInt64("id")

	result, err := h.userUseCase.UpdateUserDivisionAndRoleByUserIDUseCase(input, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}
