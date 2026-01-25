package auth

import (
	"xuetu-project/internal/model/auth"
	"xuetu-project/internal/repository"
)

type RoleService interface {
	CreateRole(role *auth.Role) error
}

type roleService struct {
	repo *repository.Repository
}

func NewRoleService(repo *repository.Repository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) CreateRole(role *auth.Role) error {
	return s.repo.RoleRepo.Create(role)
}
