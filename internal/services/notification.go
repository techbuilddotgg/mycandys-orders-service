package services

import "net/http"

type NotificationService struct {
	URL string
}

func (s *NotificationService) SendEmail() error {
	_, err := http.NewRequest("POST", s.URL, nil)
	if err != nil {
		return err
	}
	return nil
}
