package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	date := time.Date(2025, time.August, 17, 0, 0, 0, 0, time.UTC)
	dur := time.Until(date)
	return int64(dur.Hours() / 24)
}
