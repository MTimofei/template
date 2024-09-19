package main

import (
	"template/internal/controller"
	"template/internal/service"
)

func main() {
	opt := controller.Options{
		Address: "0.0.0.0:10000",
		Mode:    "http",
		Version: "v1",
	}
	s := service.NewService()
	c := controller.NewController(opt, s)
	c.Start()
}
