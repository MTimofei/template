package service

import "context"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Get(context.Context) {}
