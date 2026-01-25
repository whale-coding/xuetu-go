package controller

import (
	"xuetu-project/internal/controller/auth"
	"xuetu-project/internal/service"
)

// Controller 控制器结构体
type Controller struct {
	UserController auth.UserController
	RoleController auth.RoleController
}

// NewController 创建控制器
func NewController(svc *service.Service) *Controller {
	return &Controller{
		UserController: auth.NewUserController(svc),
		RoleController: auth.NewRoleController(svc),
	}
}
