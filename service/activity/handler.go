package activity

import (
	"absensi-karyawan-api/helper"
	"absensi-karyawan-api/service/activity/form"
	"absensi-karyawan-api/service/activity/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type activityHandler struct {
	activityUseCase usecase.ActivityUseCase
}

func NewActivityHandler(activityUseCase usecase.ActivityUseCase) *activityHandler {
	return &activityHandler{activityUseCase}
}

func (h *activityHandler) CreateActivityHandler(c *gin.Context) {
	var input form.Activity
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	id := c.GetInt64("id")

	result, err := h.activityUseCase.CreateActivityUseCase(input, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), result))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *activityHandler) UpdateActivityHandler(c *gin.Context) {
	var input form.UpdateActivity
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	id := c.GetInt64("id")

	result, err := h.activityUseCase.UpdateActivityUseCase(input, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), result))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *activityHandler) GetActivityByIDHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	result, err := h.activityUseCase.GetActivityByIDUseCase(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), result))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *activityHandler) GetAllActivityHandler(c *gin.Context) {
	result, err := h.activityUseCase.GetAllActivityUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)

}

func (h *activityHandler) DeleteActivityHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	result, err := h.activityUseCase.DeleteActivityUseCase(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), result))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)

}
