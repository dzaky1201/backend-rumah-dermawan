package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"rumahdermawan/backedn-rdi/controller/period"
	"rumahdermawan/backedn-rdi/controller/user"
	"rumahdermawan/backedn-rdi/model/web"
	user2 "rumahdermawan/backedn-rdi/service/user"
	"strings"
)

func NewRouter(userController user.UserController, periodController period.PeriodController, token user.UserToken, userService user2.UserService) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")
	api.POST("/user/register", userController.SaveUser)
	api.POST("/user/login", userController.FindByEmail)
	api.POST("/period/create", authMiddleware(token, userService), periodController.Save)

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
