package repository

import (
	"gorm.io/gorm"
)

// Repository 仓库结构体,仓库统一聚合体，业务层通过它访问各个具体的仓库
type Repository struct {
	// auth 相关的
	UserRepo UserRepository
	// subject 相关
	SubjectRepo SubjectInfoRepository
}

// NewRepository 构造函数，注入所有具体的仓库实现
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepo:    NewUserUserRepository(db),
		SubjectRepo: NewSubjectInfoRepository(db),
	}
}
