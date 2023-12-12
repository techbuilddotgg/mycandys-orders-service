package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mycandys/orders/internal/env"
	"log"
	"net/http"
)

type IAuthService interface {
	ValidateToken(token string) (*VerifyTokenResponse, error)
}

type AuthService struct {
	URL string
}

func NewAuthService() *AuthService {
	authServiceURL, err := env.GetEnvVar(env.AUTH_SERVICE_URL)
	if err != nil {
		log.Fatal(err)
	}

	return &AuthService{
		URL: authServiceURL,
	}
}

type VerifyTokenResponse struct {
	UserId string `json:"userId"`
}

func (s *AuthService) ValidateToken(token string) (*VerifyTokenResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/verify", s.URL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("unauthorized")
	}

	var response VerifyTokenResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
