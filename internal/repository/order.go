package repository

import (
	"context"
	"github.com/mycandys/orders/internal/database"
	"github.com/mycandys/orders/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type OrderRepository struct {
	coll *mongo.Collection
}

func NewOrderRepository() IOrderRepository[*models.Order, models.CreateOrderDTO, models.UpdateOrderDTO, bson.D] {
	return &OrderRepository{
		coll: database.Db.Collection("orders"),
	}
}

func (r *OrderRepository) FindOne(id string) (*models.Order, error) {
	var order models.Order

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectId}}
	err := r.coll.FindOne(context.Background(), filter).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) FindMany(filter bson.D) ([]*models.Order, error) {
	orders := make([]*models.Order, 0)

	cursor, err := r.coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil

}

func (r *OrderRepository) FindAll() ([]*models.Order, error) {
	return r.FindMany(bson.D{})
}

func (r *OrderRepository) FindByUser(id string) ([]*models.Order, error) {
	return r.FindMany(bson.D{{"user_id", id}})
}

func (r *OrderRepository) FindByStatus(status models.OrderStatus) ([]*models.Order, error) {
	return r.FindMany(bson.D{{"status", status}})
}

func (r *OrderRepository) FindByUserAndStatus(id string, status models.OrderStatus) ([]*models.Order, error) {
	return r.FindMany(bson.D{{"user_id", id}, {"status", status}})
}

func (r *OrderRepository) InsertOne(data models.CreateOrderDTO) (*models.Order, error) {
	order := models.NewOrder(data)
	_, err := r.coll.InsertOne(context.Background(), order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) UpdateOne(id string, data models.UpdateOrderDTO) (*models.Order, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", objectId}}

	update := bson.D{
		{"$set", bson.D{
			{"status", data.Status},
			{"delivered_at", data.DeliveredAt},
			{"updated_at", time.Now().Format(time.DateTime)},
		}},
	}

	var order models.Order
	err := r.coll.FindOneAndUpdate(
		context.Background(), filter, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) DeleteOne(id string) (*models.Order, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", objectId}}

	var order models.Order
	err := r.coll.FindOneAndDelete(context.Background(), filter).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) DeleteMany(filter bson.D) error {
	_, err := r.coll.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) DeleteAllByUser(userId string) error {
	return r.DeleteMany(bson.D{{"user_id", userId}})
}

func (r *OrderRepository) DeleteAll() error {
	return r.DeleteMany(bson.D{})
}
