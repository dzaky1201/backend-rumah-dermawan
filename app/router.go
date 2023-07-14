package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"rumahdermawan/backedn-rdi/controller/activity"
	"rumahdermawan/backedn-rdi/controller/period"
	"rumahdermawan/backedn-rdi/controller/user"
	"rumahdermawan/backedn-rdi/model/web"
	user2 "rumahdermawan/backedn-rdi/service/user"
	"strings"
)

func NewRouter(userController user.UserController, periodController period.PeriodController, activityController activity.ActivityController, token user.UserToken, userService user2.UserService) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")
	api.POST("/user/register", userController.SaveUser)
	api.POST("/user/login", userController.FindByEmail)

	api.POST("/period/create", authMiddleware(token, userService), periodController.Save)
	api.GET("/period/:id", authMiddleware(token, userService), periodController.FindById)
	api.PUT("/period/update/:id", authMiddleware(token, userService), periodController.Update)
	api.DELETE("/period/delete/:id", authMiddleware(token, userService), periodController.Delete)
	api.GET("/periods", authMiddleware(token, userService), periodController.FindAll)

	api.POST("/activity/create/:createType", authMiddleware(token, userService), activityController.Save)
	api.PUT("/activity/update/:updateType/:id", authMiddleware(token, userService), activityController.Update)
	api.GET("/activity/:findType/:id", authMiddleware(token, userService), activityController.FindById)
	api.DELETE("/activity/delete/:deleteType/:id", authMiddleware(token, userService), activityController.Delete)
	api.GET("/activity/list/:findAllType", authMiddleware(token, userService), activityController.FindAll)
	api.GET("/activity/report", authMiddleware(token, userService), activityController.FindReportActivity)

	return router
}

func authMiddleware(autService user.UserToken, userService user2.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := autService.ValidateToken(tokenString)
		if err != nil {
			response := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := uint(claim["user_id"].(float64))

		currentUser := userService.FindById(userId)

		c.Set("currentUser", currentUser)
	}
}
