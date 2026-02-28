package router

import (
	"github.com/MarcelloBB/ticker/docs"
	"github.com/MarcelloBB/ticker/internal/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, dbConnection *gorm.DB) {
	HealthcheckController := controller.NewHealthcheckController()

	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath

	// @BasePath /api/
	api := r.Group(basePath)
	{
		api.GET("/healthcheck", HealthcheckController.GetPing)
	}

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
