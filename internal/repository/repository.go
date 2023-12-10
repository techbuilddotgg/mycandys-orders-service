package repository

import "github.com/mycandys/orders/internal/models"

type Repository[TModel interface{}, TCreateModel interface{}, TUpdateModel interface{}, TFilter interface{}] interface {
	FindOne(id string) (TModel, error)
	FindMany(filter TFilter) ([]TModel, error)
	InsertOne(data TCreateModel) (TModel, error)
	UpdateOne(id string, data TUpdateModel) (TModel, error)
	DeleteOne(id string) (TModel, error)
	DeleteMany(filter TFilter) error
}

type IOrderRepository[TModel interface{}, TCreateModel interface{}, TUpdateModel interface{}, TFilter interface{}] interface {
	Repository[TModel, TCreateModel, TUpdateModel, TFilter]
	FindAll() ([]TModel, error)
	FindByUser(id string) ([]TModel, error)
	FindByStatus(status models.OrderStatus) ([]TModel, error)
	DeleteAllByUser(id string) error
	DeleteAll() error
	FindByUserAndStatus(id string, status models.OrderStatus) ([]TModel, error)
}
