package role

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/role/form"
	"absensi-karyawan-api/service/role/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type roleHandler struct {
	roleUseCase usecase.RoleUseCase
}

func NewRoleHandler(roleUseCase usecase.RoleUseCase) *roleHandler {
	return &roleHandler{roleUseCase}
}

func (h *roleHandler) CreateRoleHandler(c *gin.Context) {
	var input form.Role
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	id := c.GetInt64("id")

	result, err := h.roleUseCase.CreateRoleUseCase(input, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), result))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *roleHandler) GetAllRoleHandler(c *gin.Context) {
	result, err := h.roleUseCase.GetAllRoleUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)

}
