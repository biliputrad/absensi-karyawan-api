package division

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/division/form"
	"absensi-karyawan-api/service/division/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type divisionHandler struct {
	divisionUseCase usecase.DivisionUseCase
}

func NewDivisionHandler(divisionUseCase usecase.DivisionUseCase) *divisionHandler {
	return &divisionHandler{divisionUseCase}
}

func (h *divisionHandler) CreateDivisionHandler(c *gin.Context) {
	var input form.Division
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	id := c.GetInt64("id")

	result, err := h.divisionUseCase.CreateDivisionUseCase(input, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), result))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *divisionHandler) GetAllDivisionHandler(c *gin.Context) {
	result, err := h.divisionUseCase.GetAllDivisionUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)

}
