package repository

import (
	"github.com/mycandys/orders/internal/database"
	"github.com/mycandys/orders/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	coll *mongo.Collection
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		coll: database.Db.Collection("orders"),
	}
}

func (r *OrderRepository) FindOne(id string) (*models.Order, error) {
	return nil, nil
}

func (r *OrderRepository) InsertOne(data *models.CreateOrderDTO) error {
	return nil
}

func (r *OrderRepository) FindAll() ([]*models.Order, error) {
	return nil, nil
}

func (r *OrderRepository) UpdateOne(id string, data *models.UpdateOrderDTO) (*models.Order, error) {
	return nil, nil
}

func (r *OrderRepository) DeleteOne(id string) (*models.Order, error) {
	return nil, nil
}
