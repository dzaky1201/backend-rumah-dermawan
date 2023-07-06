package period

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rumahdermawan/backedn-rdi/helper"
	"rumahdermawan/backedn-rdi/model/web"
	period2 "rumahdermawan/backedn-rdi/model/web/period"
	"rumahdermawan/backedn-rdi/service/period"
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
	}

	dataPeriod := controller.periodService.Save(req)

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   dataPeriod,
	}

	c.JSON(http.StatusOK, response)
}

func (controller *PeriodControllerImpl) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *PeriodControllerImpl) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *PeriodControllerImpl) FindAll(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
