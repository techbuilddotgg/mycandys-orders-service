package mocks

import (
	"github.com/mycandys/orders/internal/models"
	"github.com/stretchr/testify/mock"
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
