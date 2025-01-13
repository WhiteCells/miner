package route

import (
	"miner/controller"

	"github.com/gin-gonic/gin"
)

type BscApiKeyRoute struct {
	bscApiKeyController *controller.BscApiKeyController
}

func NewBscApiKeyRoute() *BscApiKeyRoute {
	return &BscApiKeyRoute{
		bscApiKeyController: controller.NewBscApiController(),
	}
}

func (ar *BscApiKeyRoute) InitBscApiKeyRoute(r *gin.Engine) {
	route := r.Group("/api/bsc")
	{
		route.GET("/balance", ar.bscApiKeyController.GetTokenBalance)
	}
}
