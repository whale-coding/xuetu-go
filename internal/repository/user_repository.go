package repository

import (
	"errors"
	"xuetu-project/internal/model"
	"xuetu-project/internal/types"

	"gorm.io/gorm"
)

// UserRepository 用户仓库接口
type UserRepository interface {
	BeginTransaction() *gorm.DB                // 开启一个事务
	Transaction(func(tx *gorm.DB) error) error // 执行事务

	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
	UpdateStatus(id uint, status int) error
	Query() *gorm.DB
	QueryUsersWithCondition(query *types.UserQuery) ([]*model.User, int64, error)
}

// userRepository 用户仓库实现   UserRepo
type userRepository struct {
	db *gorm.DB
}

// NewUserUserRepository 创建用户仓库实例
func NewUserUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// BeginTransaction 开始一个新的数据库事务
func (r *userRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

// Transaction 执行事务
func (r *userRepository) Transaction(txFunc func(tx *gorm.DB) error) error {
	return r.db.Transaction(txFunc)
}

// ------------------------------------------------------------------

// Create 创建用户 ✅
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// GetByID 根据ID获取用户
func (r *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

// GetByUsername 按用户名查询用户 ✅
func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("user_name = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户
func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户（软删除）
func (r *userRepository) Delete(id uint) error {
	var user model.User
	return r.db.Delete(&user, id).Error
}

// UpdateStatus 更新用户状态
func (r *userRepository) UpdateStatus(id uint, status int) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Update("status", status).Error
}

// Query 返回查询构建器
func (r *userRepository) Query() *gorm.DB {
	return r.db.Model(&model.User{})
}

// QueryUsersWithCondition 查询用户
func (r *userRepository) QueryUsersWithCondition(query *types.UserQuery) ([]*model.User, int64, error) {
	var users []*model.User

	var count int64
	db := r.db

	if query.Username != "" {
		db = db.Where("user_name LIKE ?", "%"+query.Username+"%")
	}

	if query.Status != 0 {
		db = db.Where("status = ?", query.Status)
	}

	db = db.Order("id") // 假设按照 ID 排序
	result := db.Offset((query.Page - 1) * query.PageSize).Limit(query.PageSize).Find(&users).Count(&count)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, count, nil
}
