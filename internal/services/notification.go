package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mycandys/orders/internal/env"
	"log"
	"net/http"
)

type EmailData struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Type    string `json:"type"`
	UserID  string `json:"user"`
}

type NotificationService struct {
	URL string
}

func NewNotificationService() *NotificationService {
	notificationsServiceURL, err := env.GetEnvVar(env.NOTIFICATIONS_SERVICE_URL)
	if err != nil {
		log.Fatal(err)
	}

	return &NotificationService{
		URL: notificationsServiceURL,
	}
}

func (s *NotificationService) SendEmail(data EmailData) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/emails", s.URL), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
