package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mycandys/orders/internal/handlers"
	"github.com/mycandys/orders/internal/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	app := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	middleware := middlewares.NewMiddleware()

	app.Use(middleware.Logger())
	app.Use(cors.New(config))
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.GET("/health", handlers.HealthCheck)

	setupOrdersRoutes(app, middleware)

	return app
}
