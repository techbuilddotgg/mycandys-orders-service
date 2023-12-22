package services

import (
	"bytes"
	"fmt"
	"github.com/mycandys/orders/internal/env"
	"net/http"
)

type AnalyticsService struct {
	URL string
}

func NewAnalyticsService() *AnalyticsService {
	analyticsUrl, _ := env.GetEnvVar(env.ANALYTICS_SERVICE_URL)

	return &AnalyticsService{
		URL: analyticsUrl,
	}
}

func (s *AnalyticsService) SendEndpointCall(url string) error {

	var body = []byte(fmt.Sprintf(`{"url": "%s"}`, url))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/analytics", s.URL), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("analytics server returned status code %d", res.StatusCode)
	}

	return nil
}
