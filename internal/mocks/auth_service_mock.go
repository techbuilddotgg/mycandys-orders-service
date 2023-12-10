package mocks

import (
	"github.com/mycandys/orders/internal/services"
	"github.com/stretchr/testify/mock"
)

type AuthServiceMock struct {
	mock.Mock
}

func (_m *AuthServiceMock) ValidateToken(token string) (*services.VerifyTokenResponse, error) {
	ret := _m.Called(token)

	var r0 *services.VerifyTokenResponse
	if rf, ok := ret.Get(0).(func(string) *services.VerifyTokenResponse); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.VerifyTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
