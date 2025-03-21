package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userController         *controller.UserController
	operLogController      *controller.OperLogController
	pointsRecordController *controller.PointsRecordController
}

func NewUserRoute() *UserRoute {
	return &UserRoute{
		userController:         controller.NewUserController(),
		operLogController:      controller.NewOperLogController(),
		pointsRecordController: controller.NewPointsRecordController(),
	}
}

func (ur *UserRoute) InitUserRoute(r *gin.Engine) {
	route := r.Group("/user")
	{
		route.POST("/register", middleware.RegisterAuth(), ur.userController.Register)
		route.POST("/login", ur.userController.Login, middleware.LoginLog())
	}
	route.Use(middleware.JWTAuth())
	// route.Use(middleware.IPAuth()) // IP 验证要在 token 解析之后
	route.Use(middleware.RoleAuth(role.User, role.Admin))
	route.Use(middleware.StatusAuth())
	limiter := middleware.NewLimiter(1, time.Second)
	{
		route.POST("/logout", ur.userController.Logout)
		// route.POST("/update_passwd", ur.userController.UpdatePasswd)
		route.GET("/balance", limiter.Limit(), ur.userController.GetPointsBalance)
		route.GET("/oper_logs", ur.operLogController.GetOperLogs)
		route.GET("/points_records", ur.pointsRecordController.GetPointsRecords)
		route.GET("/address", ur.userController.GetUserAddress)
		// coin
		// route.GET("/coins", ur.userController.GetCoins)
		// pool
		// route.GET("/pools", ur.userController.GetPools)
		// soft
		// route.GET("/soft/list", ur.userController.GetSoftAll)
		//route.POST("/soft", ur.userController.SetSoft)
		//route.DELETE("/soft", ur.userController.DelSoft)
		//route.PUT("/soft", ur.userController.UpdateSoft)
		//route.GET("/soft", ur.userController.GetSoft)
		// Routers
		route.GET("/get_routers", ur.userController.GetRouters)
	}
}
