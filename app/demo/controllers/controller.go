package controllers

import (
	"github.com/jin-cg/grpc-demo/app/demo/services"
)

type Controller struct {
	s *services.Service
}

func NewController(s *services.Service) *Controller {
	return &Controller{s: s}
}
