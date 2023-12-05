package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Item struct {
	ID          string  `bson:"_id" json:"id"`
	Name        string  `bson:"name" json:"name"`
	Price       float64 `bson:"price" json:"price"`
	Description string  `bson:"description" json:"description"`
	Category    string  `bson:"category" json:"category"`
	ImageUrl    string  `bson:"image_url" json:"imgUrl"`
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	//OrderStatusCancelled OrderStatus = "cancelled"
	//OrderStatusReturned  OrderStatus = "returned"
	//OrderStatusRefunded  OrderStatus = "refunded"
	//OrderStatusFailed    OrderStatus = "failed"
	//OrderStatusPaid     OrderStatus = "paid"
)

func IsOrderStatusValid(status string) bool {
	switch status {
	case "pending", "shipped", "delivered":
		return true
	default:
		return false
	}
}

type Order struct {
	ID                   primitive.ObjectID `bson:"_id" json:"id"`
	UserID               string             `bson:"user_id" json:"userId"`
	Items                []Item             `bson:"items" json:"items"`
	Cost                 float64            `bson:"cost" json:"cost"`
	Status               OrderStatus        `bson:"status" json:"status"`
	ExpectedDeliveryDate string             `bson:"expected_delivery_date" json:"expectedDeliveryDate"`
	DeliveredAt          string             `bson:"delivered_at" json:"deliveredAt"`
	Address              string             `bson:"address" json:"address"`
	Country              string             `bson:"country" json:"country"`
	City                 string             `bson:"city" json:"city"`
	PostalCode           string             `bson:"postal_code" json:"postalCode"`
	CreatedAt            string             `bson:"created_at" json:"createdAt"`
	UpdatedAt            string             `bson:"updated_at" json:"updatedAt"`
}

func NewOrder(dto CreateOrderDTO) *Order {
	expectedDeliveryDate := time.Now().AddDate(0, 0, 7).Format(time.DateOnly)

	return &Order{
		ID:                   primitive.NewObjectID(),
		UserID:               dto.UserId,
		Items:                dto.Items,
		Cost:                 dto.Cost,
		Status:               OrderStatusPending,
		ExpectedDeliveryDate: expectedDeliveryDate,
		Address:              dto.Address,
		Country:              dto.Country,
		City:                 dto.City,
		PostalCode:           dto.PostalCode,
		CreatedAt:            time.Now().Format(time.DateTime),
	}
}

type CreateOrderDTO struct {
	UserId     string  `json:"userId"`
	Items      []Item  `json:"items"`
	Cost       float64 `json:"cost"`
	Address    string  `json:"address"`
	Country    string  `json:"country"`
	City       string  `json:"city"`
	PostalCode string  `json:"postalCode"`
}

type UpdateOrderDTO struct {
	Status      *OrderStatus `json:"status"`
	DeliveredAt *string      `json:"deliveredAt"`
}
