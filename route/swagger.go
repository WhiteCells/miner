package route

import (
	_ "miner/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerRoute struct {
}

func NewSwaggerRoute() *SwaggerRoute {
	return &SwaggerRoute{}
}

func (m *SwaggerRoute) InitSwaggerRoute(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
