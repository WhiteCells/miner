package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type AdminRoute struct {
	adminController *controller.AdminController
}

func NewAdminRoute() *AdminRoute {
	return &AdminRoute{
		adminController: controller.NewAdminController(),
	}
}

func (ar *AdminRoute) InitAdminRoute(r *gin.Engine) {
	route := r.Group("/admin")
	route.Use(middleware.JWTAuth())
	route.Use(middleware.IPAuth())
	route.Use(middleware.RoleAuth(role.Admin))
	{
		route.GET("/all_users", ar.adminController.GetAllUser)
		route.GET("/user_oper_logs", ar.adminController.GetUserOperLogs)
		route.GET("/user_login_logs", ar.adminController.GetUserLoginLogs)
		route.GET("/user_points_records", ar.adminController.GetUserPointsRecords)
		route.GET("/user_farms", ar.adminController.GetUserFarms)
		route.GET("/user_miners", ar.adminController.GetUserMiners)
		route.POST("/switch_register", ar.adminController.SwitchRegister)
		route.POST("/global_fs", ar.adminController.SetGlobalFs)
		// route.GET("/global_fs", ar.adminController.GetGlobalFs)
		route.POST("/invite_reward", ar.adminController.SetInviteReward)
		route.POST("/recharge_ratio", ar.adminController.SetRechargeRatio)
		route.POST("/user_status", ar.adminController.SetUserStatus)
		route.POST("/miner_pool_cost", ar.adminController.SetMinePoolCost)
		route.POST("/mnemonic", ar.adminController.SetMnemonic)
		route.GET("/mnemonic", ar.adminController.GetMnemonic)
		route.GET("/all_mnemonic", ar.adminController.GetAllMnemonic)
	}
}
