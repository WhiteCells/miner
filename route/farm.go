package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type FarmRoute struct {
	farmController *controller.FarmController
}

func NewFarmRoute() *FarmRoute {
	return &FarmRoute{
		farmController: controller.NewFarmController(),
	}
}

func (fr *FarmRoute) InitFarmRoute(r *gin.Engine) {
	route := r.Group("/farm")
	// 测试 limiter
	// limiter := middleware.NewLimiter(10, time.Minute) // 每个用户每分钟请求限制 10 次请求
	route.Use(middleware.JWTAuth())
	// route.Use(middleware.IPAuth())
	route.Use(middleware.RoleAuth(role.User, role.Admin))
	route.Use(middleware.StatusAuth())
	// route.Use(limiter.Limit())
	{
		route.POST("", fr.farmController.CreateFarm)
		route.DELETE("", fr.farmController.DelFarm)
		route.PUT("", fr.farmController.UpdateFarm)
		route.GET("", fr.farmController.GetFarmsByUserID)
		route.GET("/all", fr.farmController.GetFarms)
		route.GET("/:farm_id", fr.farmController.GetFarmByFarmID)
		route.PUT("/:farm_id/apply_fs", fr.farmController.ApplyFs)
		route.PUT("/transfer", fr.farmController.Transfer)
	}
}
