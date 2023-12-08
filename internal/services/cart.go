package services

import (
	"fmt"
	"github.com/mycandys/orders/internal/env"
	"log"
	"net/http"
)

type CartService struct {
	URL string
}

func NewCartService() *CartService {
	cartServiceURL, err := env.GetEnvVar(env.CART_SERVICE_URL)
	if err != nil {
		log.Fatal(err)
	}

	return &CartService{
		URL: cartServiceURL,
	}
}

func (s *CartService) ClearCart(cartId string) error {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/carts/%s/clear", s.URL, cartId), nil)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
