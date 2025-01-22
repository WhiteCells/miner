package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type WalletRoute struct {
	walletController *controller.WalletController
}

func NewWalletRoute() *WalletRoute {
	return &WalletRoute{
		walletController: controller.NewWalletController(),
	}
}

func (wr *WalletRoute) InitWalletRoute(r *gin.Engine) {
	route := r.Group("/wallet")
	route.Use(middleware.JWTAuth())
	// route.Use(middleware.IPAuth())
	route.Use(middleware.RoleAuth(role.User, role.Admin))
	route.Use(middleware.StatusAuth())
	{
		route.POST("", wr.walletController.CreateWallet)
		route.DELETE("", wr.walletController.DeleteWallet)
		route.PUT("", wr.walletController.UpdateWallet)
		route.GET("", wr.walletController.GetAllWallet)
		route.GET("/coin", wr.walletController.GetAllWalletAllCoin)
	}
}
