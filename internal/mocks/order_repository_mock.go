package mocks

import (
	"github.com/mycandys/orders/internal/models"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (_m *OrderRepositoryMock) FindAll() ([]*models.Order, error) {
	ret := _m.Called()

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func() []*models.Order); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OrderRepositoryMock) FindByUser(id string) ([]*models.Order, error) {
	ret := _m.Called(id)

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(string) []*models.Order); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OrderRepositoryMock) FindOne(id string) (*models.Order, error) {
	ret := _m.Called(id)

	var r0 *models.Order
	if rf, ok := ret.Get(0).(func(string) *models.Order); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OrderRepositoryMock) FindByStatus(status models.OrderStatus) ([]*models.Order, error) {
	ret := _m.Called(status)

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(models.OrderStatus) []*models.Order); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.OrderStatus) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OrderRepositoryMock) InsertOne(order models.CreateOrderDTO) (*models.Order, error) {
	ret := _m.Called(order)

	var r0 *models.Order
	if rf, ok := ret.Get(0).(func(models.CreateOrderDTO) *models.Order); ok {
		r0 = rf(order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.CreateOrderDTO) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OrderRepositoryMock) UpdateOne(id string, order models.UpdateOrderDTO) (*models.Order, error) {
	ret := _m.Called(id, order)

	var r0 *models.Order
	if rf, ok := ret.Get(0).(func(string, models.UpdateOrderDTO) *models.Order); ok {
		r0 = rf(id, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, models.UpdateOrderDTO) error); ok {
		r1 = rf(id, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OrderRepositoryMock) DeleteOne(id string) (*models.Order, error) {
	ret := _m.Called(id)

	var r0 *models.Order
	if rf, ok := ret.Get(0).(func(string) *models.Order); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OrderRepositoryMock) FindMany(filter bson.D) ([]*models.Order, error) {
	ret := _m.Called(filter)

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(bson.D) []*models.Order); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.D) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OrderRepositoryMock) FindByUserAndStatus(id string, status models.OrderStatus) ([]*models.Order, error) {
	ret := _m.Called(id, status)

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(string, models.OrderStatus) []*models.Order); ok {
		r0 = rf(id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, models.OrderStatus) error); ok {
		r1 = rf(id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}

func (_m *OrderRepositoryMock) DeleteMany(filter bson.D) error {
	ret := _m.Called(filter)

	var r0 error
	if rf, ok := ret.Get(0).(func(bson.D) error); ok {
		r0 = rf(filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *OrderRepositoryMock) DeleteAllByUser(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *OrderRepositoryMock) DeleteAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
