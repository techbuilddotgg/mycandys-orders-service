package repository

import "github.com/mycandys/orders/internal/models"

type Repository[T interface{}, U interface{}, V interface{}] interface {
	FindOne(id string) (T, error)
	FindAll() ([]T, error)
	InsertOne(data U) (T, error)
	UpdateOne(id string, data V) (T, error)
	DeleteOne(id string) (T, error)
}

type IOrderRepository[T interface{}, U interface{}, V interface{}] interface {
	Repository[T, U, V]
	FindByUser(id string) ([]T, error)
	FindByStatus(status models.OrderStatus) ([]T, error)
}
