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
	// route.Use(middleware.IPAuth())
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
		route.POST("/bsc_apikey", ar.adminController.AddBscApiKey)
		route.GET("/bsc_apikey", ar.adminController.GetBscApiKey)
		route.GET("/all_bsc_apikey", ar.adminController.GetAllBscApiKey)
		route.DELETE("/bsc_apikey", ar.adminController.DelBscApiKey)
		// coin
		route.POST("/coin", ar.adminController.AddCoin)
		route.PUT("/coin", ar.adminController.AddCoin)
		route.DELETE("/coin", ar.adminController.DelCoin)
		route.GET("/coin", ar.adminController.GetCoinByName)
		route.GET("/all_coin", ar.adminController.GetAllCoin)
		// pool
		route.POST("/pool", ar.adminController.AddPool)
		route.PUT("/pool", ar.adminController.AddPool)
		route.DELETE("/pool", ar.adminController.DelPool)
		route.GET("/pool", ar.adminController.GetPoolByName)
		route.GET("/all_pool", ar.adminController.GetAllPool)
		// soft
		route.POST("/soft", ar.adminController.AddSoft)
		route.PUT("/soft", ar.adminController.AddSoft)
		route.DELETE("/soft", ar.adminController.DelSoft)
		route.GET("/soft", ar.adminController.GetSoftByName)
		route.GET("/all_soft", ar.adminController.GetAllSoft)
	}
}
