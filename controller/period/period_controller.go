package period

import "github.com/gin-gonic/gin"

type PeriodController interface {
	Save(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
}
