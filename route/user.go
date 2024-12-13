package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userController         *controller.UserController
	operLogController      *controller.OperLogController
	pointsRecordController *controller.PointsRecordController
}

func NewUserRoute() *UserRoute {
	return &UserRoute{
		userController:         controller.NewUserContorller(),
		operLogController:      controller.NewOperLogController(),
		pointsRecordController: controller.NewPointRecordController(),
	}
}

func (ur *UserRoute) InitUserRoute(r *gin.Engine) {
	route := r.Group("/user")
	{
		route.POST("/register", ur.userController.Register)
		route.POST("/login", ur.userController.Login, middleware.LoginLog())
	}
	route.Use(middleware.IPVerify())
	route.Use(middleware.JWTAuth())
	route.Use(middleware.RoleAuth(role.User))
	{
		route.POST("/logout", ur.userController.Logout)
		route.GET("/balance", ur.userController.GetPointsBalance)
		route.GET("/oper_logs", ur.operLogController.GetOperLogs)
		route.GET("/points_records", ur.pointsRecordController.GetPointsRecords)
	}
}
