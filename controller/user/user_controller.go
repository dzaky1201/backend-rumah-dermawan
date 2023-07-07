package user

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	SaveUser(c *gin.Context)
	FindByEmail(c *gin.Context)
	TestHitApi(c *gin.Context)
}
