package services

import "net/http"

type CartService struct {
	URL string
}

func (s *CartService) GetCart(cartId string) error {
	_, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		return err
	}
	return nil
}
