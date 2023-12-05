package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mycandys/orders/internal/mocks"
	_ "github.com/mycandys/orders/internal/mocks"
	"github.com/mycandys/orders/internal/models"
	_ "github.com/mycandys/orders/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrdersEmptyList(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("FindAll").Return([]*models.Order{}, nil)

	server.GET("/orders", handler.GetOrders)

	req, _ := http.NewRequest("GET", "/orders", nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rec.Body.String() != "[]" {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestGetOrders(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	order := &models.Order{
		ID:                   primitive.NewObjectID(),
		UserID:               "1",
		Items:                make([]models.Item, 0),
		Cost:                 100.0,
		Status:               models.OrderStatusPending,
		ExpectedDeliveryDate: "2021-01-01",
		DeliveredAt:          "2021-01-01",
		Address:              "address",
		Country:              "country",
		City:                 "city",
		PostalCode:           "postalCode",
		CreatedAt:            "2021-01-01",
		UpdatedAt:            "2021-01-01",
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("FindAll").Return([]*models.Order{
		order,
	}, nil)

	server.GET("/orders", handler.GetOrders)

	req, _ := http.NewRequest("GET", "/orders", nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var body []models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if len(body) != 1 {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}

	if body[0].ID != order.ID {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestGetOrderByID(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	order := &models.Order{
		ID:                   primitive.NewObjectID(),
		UserID:               "1",
		Items:                make([]models.Item, 0),
		Cost:                 100.0,
		Status:               models.OrderStatusPending,
		ExpectedDeliveryDate: "2021-01-01",
		DeliveredAt:          "2021-01-01",
		Address:              "address",
		Country:              "country",
		City:                 "city",
		PostalCode:           "postalCode",
		CreatedAt:            "2021-01-01",
		UpdatedAt:            "2021-01-01",
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("FindOne", order.ID.Hex()).Return(order, nil)

	server.GET("/orders/:id", handler.GetOrder)

	req, _ := http.NewRequest("GET", "/orders/"+order.ID.Hex(), nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var body models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if body.ID != order.ID {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestGetOrderByIDNotFound(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("FindOne", "1").Return(nil, nil)

	server.GET("/orders/:id", handler.GetOrder)

	req, _ := http.NewRequest("GET", "/orders/1", nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestCreateOrder(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	dto := models.CreateOrderDTO{
		UserId:     "1",
		Items:      make([]models.Item, 0),
		Cost:       100.0,
		Address:    "address",
		Country:    "country",
		City:       "city",
		PostalCode: "postalCode",
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("InsertOne", dto).Return(models.NewOrder(dto), nil)

	server.POST("/orders", handler.CreateOrder)

	payload, _ := json.Marshal(dto)

	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(payload))

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var body models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if body.UserID != dto.UserId {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestCreateOrderInvalidPayload(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	server.POST("/orders", handler.CreateOrder)

	req, _ := http.NewRequest("POST", "/orders", nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestUpdateOrder(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	status := models.OrderStatusShipped
	dto := models.UpdateOrderDTO{
		Status:      &status,
		DeliveredAt: nil,
	}

	order := &models.Order{
		ID:                   primitive.NewObjectID(),
		UserID:               "1",
		Items:                make([]models.Item, 0),
		Cost:                 100.0,
		Status:               models.OrderStatusPending,
		ExpectedDeliveryDate: "2021-01-01",
		DeliveredAt:          "2021-01-01",
		Address:              "address",
		Country:              "country",
		City:                 "city",
		PostalCode:           "postalCode",
		CreatedAt:            "2021-01-01",
		UpdatedAt:            "2021-01-01",
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("FindOne", order.ID.Hex()).Return(order, nil)
	handler.orders.(*mocks.OrderRepositoryMock).On("UpdateOne", order.ID.Hex(), dto).Return(order, nil)

	server.PUT("/orders/:id", handler.UpdateOrder)

	payload, _ := json.Marshal(dto)

	req, _ := http.NewRequest("PUT", "/orders/"+order.ID.Hex(), bytes.NewBuffer(payload))

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var body models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if body.ID != order.ID {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestUpdateOrderInvalidPayload(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	server.PUT("/orders/:id", handler.UpdateOrder)

	req, _ := http.NewRequest("PUT", "/orders/1", nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestDeleteOrder(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	order := &models.Order{
		ID:                   primitive.NewObjectID(),
		UserID:               "1",
		Items:                make([]models.Item, 0),
		Cost:                 100.0,
		Status:               models.OrderStatusPending,
		ExpectedDeliveryDate: "2021-01-01",
		DeliveredAt:          "2021-01-01",
		Address:              "address",
		Country:              "country",
		City:                 "city",
		PostalCode:           "postalCode",
		CreatedAt:            "2021-01-01",
		UpdatedAt:            "2021-01-01",
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("DeleteOne", order.ID.Hex()).Return(order, nil)

	server.DELETE("/orders/:id", handler.DeleteOrder)

	req, _ := http.NewRequest("DELETE", "/orders/"+order.ID.Hex(), nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteOrderNotFound(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("DeleteOne", "1").Return(nil, nil)

	server.DELETE("/orders/:id", handler.DeleteOrder)

	req, _ := http.NewRequest("DELETE", "/orders/1", nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	var body models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if body.ID != primitive.NilObjectID {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}
