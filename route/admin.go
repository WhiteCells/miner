package route

import (
	"miner/controller"

	"github.com/gin-gonic/gin"
)

type AdminRoute struct {
	admineController *controller.AdminController
}

func NewAdminRoute() *AdminRoute {
	return &AdminRoute{
		admineController: controller.NewAdmineController(),
	}
}

func (ar *AdminRoute) InitAdminRoute(r *gin.Engine) {

}
