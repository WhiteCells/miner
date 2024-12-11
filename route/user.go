package route

import (
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userController *controller.UserController
}

func NewUserRoute() *UserRoute {
	return &UserRoute{
		userController: controller.NewUserContorller(),
	}
}

func (ur *UserRoute) InitUserRoute(r *gin.Engine) {
	route := r.Group("/user")
	{
		route.POST("/register", ur.userController.Register)
		route.POST("/login", ur.userController.Login, middleware.LoginLog())
	}
	// logout
	route.Use(middleware.IPVerify())
	route.Use(middleware.JWTAuth())
	{
		// route.POST("/logout", ur.userController.Logout)
	}
}
