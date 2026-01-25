package service

import (
	"xuetu-project/internal/repository"
	"xuetu-project/internal/service/auth"
)

// Service 服务层结构体
type Service struct {
	UserService auth.UserService
	RoleService auth.RoleService
}

// NewService 创建服务层
func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: auth.NewUserService(repo),
		RoleService: auth.NewRoleService(repo),
	}
}
