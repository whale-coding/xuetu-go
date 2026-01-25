package auth

import (
	"errors"
	"xuetu-project/internal/model/auth"

	"gorm.io/gorm"
)

// RoleRepository 角色仓库接口
type RoleRepository interface {
	Create(role *auth.Role) error
	GetByID(id uint) (*auth.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

// -------------------------------CRUD 方法-----------------------------------

func (r *roleRepository) Create(role *auth.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetByID(id uint) (*auth.Role, error) {
	var role auth.Role
	if err := r.db.First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("角色不存在")
		}
		return nil, err
	}
	return &role, nil
}
