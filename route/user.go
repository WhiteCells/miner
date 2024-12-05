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

func (r *UserRoute) InitUserRoute(rg *gin.Engine) {
	publicGroup := rg.Group("/user")
	publicGroup.Use(middleware.OperLogger())
	{
		publicGroup.POST("/register", r.userController.Register)
		publicGroup.POST("/login", r.userController.Login, middleware.LoginLog())
	}
	authGroup := rg.Group("/user")
	authGroup.Use(middleware.JWTAuth())
	authGroup.Use(middleware.IPVerify())
	// authGroup.Use(middleware.RoleAuth(""))
	{
		// authGroup.PUT("/:id", r.userController)
	}
}
