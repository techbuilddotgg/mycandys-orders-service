package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mycandys/orders/internal/handlers"
)

func setupOrdersRoutes(app *gin.Engine) {
	orders := app.Group("/orders")
	ordersHandler := handlers.NewOrderHandler()

	// orders.Use(middlewares.Auth())

	orders.GET(":id", ordersHandler.GetOrder)
	orders.GET("", ordersHandler.GetOrders)
	orders.GET("/user/:id", ordersHandler.GetOrdersByUser)
	orders.GET("/status/:status", ordersHandler.GetOrderByStatus)
	orders.POST("", ordersHandler.CreateOrder)
	orders.PUT(":id", ordersHandler.UpdateOrder)
	orders.DELETE(":id", ordersHandler.DeleteOrder)
}
