package service

import (
	"xuetu-project/internal/repository"
)

// Service 服务层结构体
type Service struct {
	UserService UserService
}

// NewService 创建服务层
func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repo),
	}
}
