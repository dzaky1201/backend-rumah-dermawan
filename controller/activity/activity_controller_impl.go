package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rumahdermawan/backedn-rdi/helper"
	"rumahdermawan/backedn-rdi/model/web"
	activities2 "rumahdermawan/backedn-rdi/model/web/activities"
	"rumahdermawan/backedn-rdi/service/activities"
	"strconv"
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
	var req activities2.ActivityCreateRequest
	pathId, _ := strconv.Atoi(c.Param("id"))
	updateType := c.Param("updateType")

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

	dataUpdate, errData := controller.service.Update(req, pathId, updateType)

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
		Data:   dataUpdate,
	}

	c.JSON(http.StatusOK, response)
}

func (controller *ActivityControllerImpl) Delete(c *gin.Context) {
	pathId, _ := strconv.Atoi(c.Param("id"))
	deleteType := c.Param("deleteType")
	err := controller.service.Delete(pathId, deleteType)
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

func (controller *ActivityControllerImpl) FindAll(c *gin.Context) {
	findAllType := c.Param("findAllType")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	description := c.Query("description")

	param := activities2.ActivityQueryParam{Page: page, Limit: limit, Description: description}

	dataList, count, offset, err := controller.service.FindAll(param, findAllType)
	if err != nil {
		response := web.WebResponseList{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Count:  0,
			Offset: 0,
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := web.WebResponseList{
		Code:   http.StatusOK,
		Status: "Success",
		Count:  count,
		Offset: offset,
		Data:   dataList,
	}
	c.JSON(http.StatusOK, response)
	return

}

func (controller *ActivityControllerImpl) FindById(c *gin.Context) {
	// path by url
	pathId, _ := strconv.Atoi(c.Param("id"))
	findType := c.Param("findType")

	getDetail, errData := controller.service.FindById(pathId, findType)

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

func (controller *ActivityControllerImpl) FindReportActivity(c *gin.Context) {
	year := c.Query("year")

	data, err := controller.service.ReportActivityAll(year)
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
		Data:   data,
	}
	c.JSON(http.StatusOK, response)
	return

}
