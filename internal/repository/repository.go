package repository

import (
	"xuetu-project/internal/repository/auth"

	"gorm.io/gorm"
)

// Repository 仓库结构体,仓库统一聚合体，业务层通过它访问各个具体的仓库
type Repository struct {
	// auth 相关的
	UserRepo auth.UserRepository
	RoleRepo auth.RoleRepository
	//Subject    SubjectRepository
	//Category   CategoryRepository
	//Label      LabelRepository
	//Liked      LikedRepository
	//Role       RoleRepository
	//Permission PermissionRepository
	//Circle     CircleRepository
}

// NewRepository 构造函数，注入所有具体的仓库实现
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepo: auth.NewUserUserRepository(db),
		RoleRepo: auth.NewRoleRepository(db),
		//Subject:    NewSubjectRepository(db),
		//Category:   NewCategoryRepository(db),
		//Label:      NewLabelRepository(db),
		//Liked:      NewLikedRepository(db),
		//Role:       NewRoleRepository(db),
		//Permission: NewPermissionRepository(db),
		//Circle:     NewCircleRepository(db),
	}
}
