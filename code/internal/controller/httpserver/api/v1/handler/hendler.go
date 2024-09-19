package handler

import (
	"template/internal/controller/contract"
)

type handler struct {
	service contract.Service
}

func NewHandler(s contract.Service) *handler {
	return &handler{
		service: s,
	}
}
