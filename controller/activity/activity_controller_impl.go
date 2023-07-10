package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rumahdermawan/backedn-rdi/helper"
	"rumahdermawan/backedn-rdi/model/web"
	activities2 "rumahdermawan/backedn-rdi/model/web/activities"
	"rumahdermawan/backedn-rdi/service/activities"
)

type ActivityControllerImpl struct {
	service activities.ActivityService
}

func NewActivityController(service activities.ActivityService) *ActivityControllerImpl {
	return &ActivityControllerImpl{service: service}
}

func (controller *ActivityControllerImpl) Save(c *gin.Context) {
	var req activities2.ActivityCreateRequest

	param := c.Param("createType")

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

	dataPeriod, errData := controller.service.Save(req, param)

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

func (controller *ActivityControllerImpl) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *ActivityControllerImpl) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *ActivityControllerImpl) FindAll(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *ActivityControllerImpl) FindById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
