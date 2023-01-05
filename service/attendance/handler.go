package activity

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/attendance/form"
	"absensi-karyawan-api/service/attendance/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type attendanceHandler struct {
	attendanceUseCase usecase.AttendanceUseCase
}

func NewAttendanceHandler(attendanceUseCase usecase.AttendanceUseCase) *attendanceHandler {
	return &attendanceHandler{attendanceUseCase}
}

func (h *attendanceHandler) CreateAttedanceCheckInHandler(c *gin.Context) {
	id := c.GetInt64("id")

	result, err := h.attendanceUseCase.CheckInAttendanceUseCase(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), result))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *attendanceHandler) CreateAttedanceCheckOutHandler(c *gin.Context) {
	id := c.GetInt64("id")

	result, err := h.attendanceUseCase.CheckOutAttendanceUseCase(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), result))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *attendanceHandler) GetAttendanceByIDHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	result, err := h.attendanceUseCase.GetAttendanceByIDUseCase(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *attendanceHandler) GetAllAttendanceHandler(c *gin.Context) {
	id := c.GetInt64("id")

	result, err := h.attendanceUseCase.GetAllAttendanceUseCase(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *attendanceHandler) GetAttendanceByDateHandler(c *gin.Context) {
	var input form.AttedanceDate
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	id := c.GetInt64("id")

	result, err := h.attendanceUseCase.GetAttendanceByDateUseCase(input, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)

}
