package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mycandys/orders/internal/handlers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.GET("/health", handlers.HealthCheck)

	return app
}
