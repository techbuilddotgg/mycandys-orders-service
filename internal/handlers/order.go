package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mycandys/orders/internal/models"
	"github.com/mycandys/orders/internal/repository"
	"github.com/mycandys/orders/internal/services"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type OrderHandler struct {
	orders        repository.IOrderRepository[*models.Order, models.CreateOrderDTO, models.UpdateOrderDTO, bson.D]
	notifications *services.NotificationService
	carts         *services.CartService
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		orders:        repository.NewOrderRepository(),
		notifications: services.NewNotificationService(),
		carts:         services.NewCartService(),
	}
}

// GetOrder Order godoc
// @Summary get order by id
// @Tags orders
// @Schemes
// @Description get order by id
// @Param id path string true "order id"
// @Success 200
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	o, err := h.orders.FindOne(id)
	if err != nil || o == nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(200, o)
}

// GetOrders Orders godoc
// @Summary get all orders
// @Tags orders
// @Schemes
// @Description get all orders
// @Success 200
// @Router /orders [get]
func (h *OrderHandler) GetOrders(c *gin.Context) {
	orders, err := h.orders.FindAll()
	if err != nil {
		c.JSON(404, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(200, orders)
}

// GetOrdersByUser Orders godoc
// @Summary get all orders by user
// @Tags orders
// @Schemes
// @Description get all orders by user
// @Param id path string true "user id"
// @Success 200
// @Router /orders/user/{id} [get]
func (h *OrderHandler) GetOrdersByUser(c *gin.Context) {
	id := c.Param("id")

	orders, err := h.orders.FindByUser(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(200, orders)
}

// GetOrderByStatus Orders godoc
// @Summary get all orders by status
// @Tags orders
// @Schemes
// @Description get all orders by status
// @Param status path string true "order status"
// @Success 200
// @Router /orders/status/{status} [get]
func (h *OrderHandler) GetOrderByStatus(c *gin.Context) {
	status := c.Param("status")

	isValid := models.IsOrderStatusValid(status)
	if !isValid {
		c.JSON(400, gin.H{"error": "Invalid order status"})
		return
	}

	orders, err := h.orders.FindByStatus(models.OrderStatus(status))
	if err != nil {
		c.JSON(404, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(200, orders)
}

// CreateOrder Order godoc
// @Summary create order
// @Tags orders
// @Schemes
// @Description create order
// @Accept json
// @Produce json
// @Param order body models.CreateOrderDTO true "order"
// @Success 201
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var dto models.CreateOrderDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orders.InsertOne(dto)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cloud not create order"})
		return
	}

	if h.carts != nil {
		err = h.carts.ClearCart(dto.CartID)
		if err != nil {
			log.Print(err.Error())
		}
	}

	if h.notifications != nil {
		err = h.notifications.SendEmail(services.NewOrderCreatedEmail(order.UserID, order.ID.String()))
		if err != nil {
			log.Print(err.Error())
		}
	}

	c.JSON(201, order)
}

// UpdateOrder Order godoc
// @Summary update order
// @Tags orders
// @Schemes
// @Description update order
// @Accept json
// @Produce json
// @Param id path string true "order id"
// @Param order body models.UpdateOrderDTO true "order"
// @Success 200
// @Router /orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var dto models.UpdateOrderDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if dto.DeliveredAt != nil {
		_, err := time.Parse(time.DateTime, *dto.DeliveredAt)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid date time format"})
			return
		}
	}

	if *dto.Status == models.OrderStatusDelivered && dto.DeliveredAt == nil {
		dto.DeliveredAt = new(string)
		*dto.DeliveredAt = time.Now().Format(time.DateTime)
	}

	order, err := h.orders.UpdateOne(id, dto)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cloud not update order"})
		return
	}

	if h.notifications != nil {
		err = h.notifications.SendEmail(
			services.NewOrderStatusUpdatedEmail(
				order.UserID, order.ID.String(), order.Status))
		if err != nil {
			log.Print(err.Error())
		}
	}

	c.JSON(200, order)
}

// DeleteOrder Order godoc
// @Summary delete order
// @Tags orders
// @Schemes
// @Description delete order
// @Param id path string true "order id"
// @Success 200
// @Router /orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.orders.DeleteOne(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cloud not delete order"})
		return
	}

	c.JSON(200, order)
}

// GetMyOrders Orders godoc
// @Summary get all orders by user
// @Tags orders
// @Schemes
// @Description get all orders by user
// @Security ApiKeyAuth
// @Success 200
// @Router /orders/me [get]
func (h *OrderHandler) GetMyOrders(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	orders, err := h.orders.FindByUser(userId)
	if err != nil {
		c.JSON(404, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(200, orders)
}

// GetMyOrdersByStatus Orders godoc
// @Summary get all orders by status
// @Tags orders
// @Schemes
// @Description get all orders by status
// @Security ApiKeyAuth
// @Param status path string true "order status"
// @Success 200
// @Router /orders/me/status/{status} [get]
func (h *OrderHandler) GetMyOrdersByStatus(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	status := c.Param("status")

	isValid := models.IsOrderStatusValid(status)
	if !isValid {
		c.JSON(400, gin.H{"error": "Invalid order status"})
		return
	}

	orders, err := h.orders.FindByUserAndStatus(userId, models.OrderStatus(status))
	if err != nil {
		c.JSON(404, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(200, orders)
}

// DeleteAllOrders Orders godoc
// @Summary delete all orders
// @Tags orders
// @Schemes
// @Description delete all orders
// @Success 200
// @Router /orders [delete]
func (h *OrderHandler) DeleteAllOrders(c *gin.Context) {
	err := h.orders.DeleteAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "Cloud not delete orders"})
		return
	}

	c.JSON(200, gin.H{"message": "All orders deleted"})
}

// DeleteAllMyOrders Orders godoc
// @Summary delete all orders by user
// @Tags orders
// @Schemes
// @Description delete all orders by user
// @Security ApiKeyAuth
// @Success 200
// @Router /orders/me [delete]
func (h *OrderHandler) DeleteAllMyOrders(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	err := h.orders.DeleteAllByUser(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cloud not delete orders"})
		return
	}

	c.JSON(200, gin.H{"message": "All orders deleted"})
}
