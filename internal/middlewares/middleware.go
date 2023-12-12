package middlewares

import (
	"github.com/mycandys/orders/internal/services"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	AuthService services.IAuthService
	logger      *logrus.Logger
}

func NewMiddleware() *Middleware {
	return &Middleware{
		AuthService: services.NewAuthService(),
		logger:      logrus.New(),
	}
}
