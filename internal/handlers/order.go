package handlers

import (
	"github.com/mycandys/orders/internal/models"
	"github.com/mycandys/orders/internal/repository"
)

// CRUD operations for orders
// ADD Payment operations?

type OrderHandler struct {
	orders repository.Repository[*models.Order, *models.CreateOrderDTO, *models.UpdateOrderDTO]
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		orders: repository.NewOrderRepository(),
	}
}
