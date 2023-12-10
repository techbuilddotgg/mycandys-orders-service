package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mycandys/orders/internal/handlers"
	"github.com/mycandys/orders/internal/middlewares"
)

func setupOrdersRoutes(app *gin.Engine) {
	orders := app.Group("/orders")
	ordersHandler := handlers.NewOrderHandler()

	middleware := middlewares.NewMiddleware()
	requiredAuth := orders.Use(middleware.Auth())

	orders.GET(":id", ordersHandler.GetOrder)
	orders.GET("", ordersHandler.GetOrders)
	orders.GET("/user/:id", ordersHandler.GetOrdersByUser)
	orders.GET("/status/:status", ordersHandler.GetOrderByStatus)
	orders.POST("", ordersHandler.CreateOrder)
	orders.PUT(":id", ordersHandler.UpdateOrder)
	orders.DELETE(":id", ordersHandler.DeleteOrder)
	orders.DELETE("", ordersHandler.DeleteAllOrders)

	requiredAuth.GET("/me", ordersHandler.GetMyOrders)
	requiredAuth.GET("/me/status/:status", ordersHandler.GetMyOrdersByStatus)
	requiredAuth.DELETE("/me", ordersHandler.DeleteAllMyOrders)
}
