package router

import (
	"context"
	"time"

	"github.com/MarcelloBB/ticker/docs"
	"github.com/MarcelloBB/ticker/internal/config"
	"github.com/MarcelloBB/ticker/internal/controller"
	"github.com/MarcelloBB/ticker/internal/repository"
	"github.com/MarcelloBB/ticker/internal/service"
	"github.com/MarcelloBB/ticker/internal/worker"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, dbConnection *gorm.DB) {
	HealthcheckController := controller.NewHealthcheckController()
	uptimeRepository := repository.NewUptimeGormRepository(dbConnection)
	uptimeService := service.NewUptimeService(uptimeRepository)
	uptimeController := controller.NewUptimeController(uptimeService)
	checkInterval := config.LoadConfigIni("uptime", "check_interval_seconds", 30).(int)
	checkConcurrency := config.LoadConfigIni("uptime", "check_concurrency", 5).(int)
	uptimeWorker := worker.NewUptimeWorker(
		uptimeService,
		time.Duration(checkInterval)*time.Second,
		checkConcurrency,
	)
	uptimeWorker.Start(context.Background())

	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath

	// @BasePath /api/
	api := r.Group(basePath)
	{
		api.GET("/healthcheck", HealthcheckController.GetPing)
		api.POST("/uptime/targets", uptimeController.CreateTarget)
		api.GET("/uptime/targets", uptimeController.ListTargets)
	}

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
