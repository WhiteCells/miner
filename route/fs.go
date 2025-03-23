package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type FsRoute struct {
	fsController *controller.FsController
}

func NewFsRoute() *FsRoute {
	return &FsRoute{
		fsController: controller.NewFsController(),
	}
}

func (fr *FsRoute) InitFsRoute(r *gin.Engine) {
	route := r.Group("/fs")
	route.Use(middleware.JWTAuth())
	// route.Use(middleware.IPAuth())
	route.Use(middleware.RoleAuth(role.User, role.Admin))
	route.Use(middleware.StatusAuth())
	{
		route.POST("", fr.fsController.CreateFs)
		route.DELETE("", fr.fsController.DeleteFs)
		route.PUT("", fr.fsController.UpdateFs)
		route.GET("", fr.fsController.GetFss)
		route.GET("/:fs_id", fr.fsController.GetFsByID)
		// route.PUT("/apply_wallet", fr.fsController.ApplyWallet)
	}
}
