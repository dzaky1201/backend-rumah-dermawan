package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rumahdermawan/backedn-rdi/helper"
	"rumahdermawan/backedn-rdi/model/web"
	user2 "rumahdermawan/backedn-rdi/model/web/user"
	"rumahdermawan/backedn-rdi/service/user"
)

type UserControllerImpl struct {
	userService user.UserService
	userToken   UserToken
}

func NewUserController(userService user.UserService, userToken UserToken) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userService,
		userToken:   userToken,
	}
}

func (controller *UserControllerImpl) SaveUser(c *gin.Context) {
	var request user2.UserCreateRequest

	errInput := c.ShouldBindJSON(&request)
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

	dataRegister := controller.userService.SaveUser(request, "")

	tokenUser, errToken := controller.userToken.GenerateToken(dataRegister.Id)

	if errToken != nil {
		response := web.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: "Register account failed",
			Data:   nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataRegister.Token = tokenUser

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Account has been register",
		Data:   dataRegister,
	}

	c.JSON(http.StatusOK, response)

}

func (controller *UserControllerImpl) FindByEmail(c *gin.Context) {
	var request user2.UserLoginRequest

	errInput := c.ShouldBindJSON(&request)
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

	dataLogin := controller.userService.FindByEmail(request, "")

	tokenUser, errToken := controller.userToken.GenerateToken(dataLogin.Id)

	if errToken != nil {
		response := web.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: "Register account failed",
			Data:   nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataLogin.Token = tokenUser

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Account has been register",
		Data:   dataLogin,
	}

	c.JSON(http.StatusOK, response)
}

func (controller *UserControllerImpl) TestHitApi(c *gin.Context) {
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success hit api",
		Data:   nil,
	}

	c.JSON(http.StatusOK, response)
}
