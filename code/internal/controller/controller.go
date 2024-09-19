package controller

import (
	"net/http"
	"template/internal/controller/contract"
	"template/internal/controller/httpserver"
	"template/internal/controller/httpserver/api/v1/handler"
	"template/internal/controller/httpserver/api/v1/router"
)

type Options struct {
	Address string
	Mode    string
	Version string
}

type controller interface {
	Start() error
	Stop() error
}

type Controller struct {
	c controller
}

func NewController(o Options, s contract.Service) *Controller {
	var c controller
	switch o.Mode {
	case "http":
		var r *http.ServeMux
		switch o.Version {
		case "v1":
			h := handler.NewHandler(s)
			r = router.Router(h)
		default:
			panic("Что-то непонятное. При создании контроллера. Версия не правальная.")
		}
		c = httpserver.NewServer(o.Address, r)
	case "rpc":
		panic("Что-то непонятное. При создании контроллера. Мод не реализован.")
	case "queue":
		panic("Что-то непонятное. При создании контроллера. Мод не реализован.")
	default:
		panic("Что-то непонятное. При создании контроллера. Мод не правальный.")
	}

	return &Controller{c}
}

func (c *Controller) Start() error {
	return c.c.Start()
}

func (c *Controller) Stop() error {
	return c.c.Stop()
}
