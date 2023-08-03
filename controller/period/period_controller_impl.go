package period

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rumahdermawan/backedn-rdi/helper"
	"rumahdermawan/backedn-rdi/model/domain"
	"rumahdermawan/backedn-rdi/model/web"
	period2 "rumahdermawan/backedn-rdi/model/web/period"
	"rumahdermawan/backedn-rdi/service/period"
	"strconv"
)

type PeriodControllerImpl struct {
	periodService period.PeriodService
}

func NewPeriodController(periodService period.PeriodService) *PeriodControllerImpl {
	return &PeriodControllerImpl{periodService}
}

func (controller *PeriodControllerImpl) Save(c *gin.Context) {
	var req period2.PeriodCreateRequest

	errInput := c.ShouldBindJSON(&req)
	if errInput != nil {
		errors := helper.FormatValidationError(errInput)
		errorMessage := gin.H{"errors": errors}
		response := web.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: "error",
			Data:   errorMessage,
		}
		c.JSON(http.StatusMethodNotAllowed, response)
		return
	}

	dataPeriod, errData := controller.periodService.Save(req)

	if errData != nil {
		response := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: errData.Error(),
			Data:   "periode sudah ada",
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   dataPeriod,
	}

	c.JSON(http.StatusOK, response)
}

func (controller *PeriodControllerImpl) Update(c *gin.Context) {
	var req period2.PeriodCreateRequest
	pathId := c.Param("id")

	errInput := c.ShouldBindJSON(&req)
	if errInput != nil {
		errors := helper.FormatValidationError(errInput)
		errorMessage := gin.H{"errors": errors}
		response := web.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: "error",
			Data:   errorMessage,
		}
		c.JSON(http.StatusMethodNotAllowed, response)
		return
	}

	id, err := strconv.ParseUint(pathId, 10, 32)
	if err != nil {
		return
	}

	requestId := domain.YearPeriod{Id: uint(id)}

	dataPeriod, errData := controller.periodService.Update(req, requestId)

	if errData != nil {
		response := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: errData.Error(),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   dataPeriod,
	}

	c.JSON(http.StatusOK, response)
}

func (controller *PeriodControllerImpl) Delete(c *gin.Context) {
	path, _ := strconv.Atoi(c.Param("id"))
	err := controller.periodService.Delete(path)
	if err != nil {
		response := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "delete Success",
		Data:   nil,
	}

	c.JSON(http.StatusOK, response)

}

func (controller *PeriodControllerImpl) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	year := c.Query("year")

	dataList, err := controller.periodService.FindAll(page, limit, year)
	if err != nil {
		response := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   dataList,
	}

	c.JSON(http.StatusOK, response)
}

func (controller *PeriodControllerImpl) FindById(c *gin.Context) {
	// path by url
	path := c.Param("id")

	// convert path to integer
	id, err := strconv.ParseUint(path, 10, 32)
	if err != nil {
		return
	}

	request := domain.YearPeriod{Id: uint(id)}
	getDetail, errData := controller.periodService.FindById(request)

	if errData != nil {
		response := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: errData.Error(),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   getDetail,
	}

	c.JSON(http.StatusOK, response)
}
