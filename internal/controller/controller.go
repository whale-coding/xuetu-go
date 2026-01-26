package controller

import (
	"xuetu-project/internal/service"
)

// Controller 控制器结构体
type Controller struct {
	UserController UserController
}

// NewController 创建控制器
func NewController(svc *service.Service) *Controller {
	return &Controller{
		UserController: NewUserController(svc),
	}
}
