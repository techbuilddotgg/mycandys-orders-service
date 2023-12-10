package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mycandys/orders/internal/middlewares"
	"github.com/mycandys/orders/internal/mocks"
	_ "github.com/mycandys/orders/internal/mocks"
	"github.com/mycandys/orders/internal/models"
	_ "github.com/mycandys/orders/internal/models"
	_ "github.com/mycandys/orders/internal/repository"
	"github.com/mycandys/orders/internal/services"
	_ "go.mongodb.org/mongo-driver/bson"
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
		orders:        &mocks.OrderRepositoryMock{},
		notifications: nil,
		carts:         nil,
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
		orders:        &mocks.OrderRepositoryMock{},
		notifications: nil,
		carts:         nil,
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

func TestGetOrdersByUser(t *testing.T) {
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

	handler.orders.(*mocks.OrderRepositoryMock).On("FindByUser", order.UserID).Return([]*models.Order{
		order,
	}, nil)

	server.GET("/orders/user/:id", handler.GetOrdersByUser)

	req, _ := http.NewRequest("GET", "/orders/user/"+order.UserID, nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	var body []models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if len(body) != 1 {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}

	if body[0].ID != order.ID {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestGetOrdersByUserNotFound(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("FindByUser", "1").Return(nil, nil)

	server.GET("/orders/user/:id", handler.GetOrdersByUser)

	req, _ := http.NewRequest("GET", "/orders/user/1", nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	var body []models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if len(body) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestGetOrderByStatus(t *testing.T) {
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

	handler.orders.(*mocks.OrderRepositoryMock).On("FindByStatus", order.Status).Return([]*models.Order{
		order,
	}, nil)

	server.GET("/orders/status/:status", handler.GetOrderByStatus)

	req, _ := http.NewRequest("GET", "/orders/status/"+string(order.Status), nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	var body []models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if len(body) != 1 {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}

	if body[0].ID != order.ID {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestGetOrderByStatusNotFound(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("FindByStatus", models.OrderStatus("invalid")).Return(nil, nil)

	server.GET("/orders/status/:status", handler.GetOrderByStatus)

	req, _ := http.NewRequest("GET", "/orders/status/invalid", nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	var body []models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if len(body) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestGetOrdersByUserAndStatus(t *testing.T) {
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

	handler.orders.(*mocks.OrderRepositoryMock).On("FindByUserAndStatus", order.UserID, order.Status).Return([]*models.Order{
		order,
	}, nil)

	middleware := &middlewares.Middleware{
		AuthService: &mocks.AuthServiceMock{},
	}

	middleware.AuthService.(*mocks.AuthServiceMock).On("ValidateToken", "token").Return(&services.VerifyTokenResponse{
		UserId: order.UserID,
	}, nil)

	requiredAuth := server.Use(middleware.Auth())

	requiredAuth.GET("/orders/me/status/:status", handler.GetMyOrdersByStatus)

	req, _ := http.NewRequest("GET", "/orders/me/status/"+string(order.Status), nil)

	req.Header.Set("Authorization", "Bearer token")

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
}

func TestGetOrdersByUserAndStatusNotFound(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	middleware := &middlewares.Middleware{
		AuthService: &mocks.AuthServiceMock{},
	}

	middleware.AuthService.(*mocks.AuthServiceMock).On("ValidateToken", "token").Return(&services.VerifyTokenResponse{
		UserId: "1",
	}, nil)

	handler.orders.(*mocks.OrderRepositoryMock).On("FindByUserAndStatus", "1", models.OrderStatus("invalid")).Return(nil, nil)

	requiredAuth := server.Use(middleware.Auth())

	requiredAuth.GET("/orders/me/status/:status", handler.GetMyOrdersByStatus)

	req, _ := http.NewRequest("GET", "/orders/me/status/invalid", nil)

	req.Header.Set("Authorization", "Bearer token")

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	var body []models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if len(body) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestGetOrdersByUserAndStatusUnauthorized(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	middleware := &middlewares.Middleware{
		AuthService: &mocks.AuthServiceMock{},
	}

	middleware.AuthService.(*mocks.AuthServiceMock).On("ValidateToken", "token").Return(nil, errors.New("unauthorized"))

	handler.orders.(*mocks.OrderRepositoryMock).On("FindByUserAndStatus", "1", models.OrderStatus("invalid")).Return(nil, nil)

	requiredAuth := server.Use(middleware.Auth())

	requiredAuth.GET("/orders/me/status/:status", handler.GetMyOrdersByStatus)

	req, _ := http.NewRequest("GET", "/orders/me/status/invalid", nil)

	req.Header.Set("Authorization", "Bearer token")

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}

func TestGetMyOrders(t *testing.T) {
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

	handler.orders.(*mocks.OrderRepositoryMock).On("FindByUser", order.UserID).Return([]*models.Order{
		order,
	}, nil)

	middleware := &middlewares.Middleware{
		AuthService: &mocks.AuthServiceMock{},
	}

	middleware.AuthService.(*mocks.AuthServiceMock).On("ValidateToken", "token").Return(&services.VerifyTokenResponse{
		UserId: order.UserID,
	}, nil)

	requiredAuth := server.Use(middleware.Auth())

	requiredAuth.GET("/orders/me", handler.GetMyOrders)

	req, _ := http.NewRequest("GET", "/orders/me", nil)

	req.Header.Set("Authorization", "Bearer token")

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
}

func TestGetMyOrdersNotFound(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	middleware := &middlewares.Middleware{
		AuthService: &mocks.AuthServiceMock{},
	}

	middleware.AuthService.(*mocks.AuthServiceMock).On("ValidateToken", "token").Return(&services.VerifyTokenResponse{
		UserId: "1",
	}, nil)

	handler.orders.(*mocks.OrderRepositoryMock).On("FindByUser", "1").Return(nil, nil)

	requiredAuth := server.Use(middleware.Auth())

	requiredAuth.GET("/orders/me", handler.GetMyOrders)

	req, _ := http.NewRequest("GET", "/orders/me", nil)

	req.Header.Set("Authorization", "Bearer token")

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	var body []models.Order
	_ = json.Unmarshal(rec.Body.Bytes(), &body)

	if len(body) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), "[]")
	}
}

func TestGetMyOrdersUnauthorized(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	middleware := &middlewares.Middleware{
		AuthService: &mocks.AuthServiceMock{},
	}

	middleware.AuthService.(*mocks.AuthServiceMock).On("ValidateToken", "token").Return(nil, errors.New("unauthorized"))

	handler.orders.(*mocks.OrderRepositoryMock).On("FindByUser", "1").Return(nil, nil)

	requiredAuth := server.Use(middleware.Auth())

	requiredAuth.GET("/orders/me", handler.GetMyOrders)

	req, _ := http.NewRequest("GET", "/orders/me", nil)

	req.Header.Set("Authorization", "Bearer token")

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}

func TestDeleteAllOrders(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("DeleteAll").Return(nil)

	server.DELETE("/orders", handler.DeleteAllOrders)

	req, _ := http.NewRequest("DELETE", "/orders", nil)

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteAllMyOrders(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("DeleteAllByUser", "1").Return(nil)

	middleware := &middlewares.Middleware{
		AuthService: &mocks.AuthServiceMock{},
	}

	middleware.AuthService.(*mocks.AuthServiceMock).On("ValidateToken", "token").Return(&services.VerifyTokenResponse{
		UserId: "1",
	}, nil)

	requiredAuth := server.Use(middleware.Auth())

	requiredAuth.DELETE("/orders/me", handler.DeleteAllMyOrders)

	req, _ := http.NewRequest("DELETE", "/orders/me", nil)

	req.Header.Set("Authorization", "Bearer token")

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteAllMyOrdersUnauthorized(t *testing.T) {
	server := gin.Default()

	handler := &OrderHandler{
		orders: &mocks.OrderRepositoryMock{},
	}

	handler.orders.(*mocks.OrderRepositoryMock).On("DeleteByUser", "1").Return(nil)

	middleware := &middlewares.Middleware{
		AuthService: &mocks.AuthServiceMock{},
	}

	middleware.AuthService.(*mocks.AuthServiceMock).On("ValidateToken", "token").Return(nil, errors.New("unauthorized"))

	requiredAuth := server.Use(middleware.Auth())

	requiredAuth.DELETE("/orders/me", handler.DeleteAllMyOrders)

	req, _ := http.NewRequest("DELETE", "/orders/me", nil)

	req.Header.Set("Authorization", "Bearer token")

	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}
