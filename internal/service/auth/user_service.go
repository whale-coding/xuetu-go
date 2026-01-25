package auth

import (
	"fmt"
	"log"
	"xuetu-project/internal/model/auth"
	"xuetu-project/internal/pkg/jwt"
	"xuetu-project/internal/repository"
	"xuetu-project/internal/types"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 用户服务接口
type UserService interface {
	Register(req *types.UserRegisterRequest) error
	Login(req *types.UserLoginRequest) (string, error)
	GetUserByUsername(username string) (*auth.User, error)
	CreateUser(user *auth.User) error
	GetUserByID(id uint) (*auth.User, error)
	UpdateUser(user *auth.User) error
	DeleteUser(id uint) error
	UpdateUserStatus(id uint, status int) error
	QueryUsers(query *types.UserQueryRequest) (*types.UserListResponse, error)
	QueryUsersWithCondition(query *types.UserQuery) (*types.PaginationResponse, error)
	BatchDeleteUser(ids []uint) error
	BatchDeleteUser2(ids []uint) error
}

// userService 用户服务实现
type userService struct {
	repo *repository.Repository
}

// NewUserService 创建用户服务实例
func NewUserService(repo *repository.Repository) UserService {
	return &userService{repo: repo}
}

// --------------------------------------------------------------

// Register 用户注册 ✅
func (s *userService) Register(req *types.UserRegisterRequest) error {
	// 检查用户名是否已存在
	existingUser, err := s.repo.UserRepo.GetByUsername(req.Username)
	if err != nil {
		return fmt.Errorf("检查用户名时出错: %v", err)
	}
	if existingUser != nil {
		return fmt.Errorf("用户名 %s 已存在", req.Username)
	}

	// 对密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("密码哈希处理失败: %v", err)
	}
	// 创建用户模型
	user := &auth.User{
		Username:  req.Username,
		Password:  string(hashedPassword), // 存储哈希后的密码
		Nickname:  req.Nickname,
		Email:     req.Email,
		Phone:     req.Phone,
		Sex:       req.Sex,
		Avatar:    req.Avatar,
		Introduce: req.Introduce,
		ExtJson:   req.ExtJson,
		Status:    1, // 默认状态为正常
	}

	// 调用数据访问层创建用户
	if err := s.repo.UserRepo.Create(user); err != nil {
		return fmt.Errorf("用户创建失败: %v", err)
	}

	log.Printf("注册请求成功，用户ID: %d, 用户名: %s\n", user.ID, user.Username)

	return nil
}

// Login 用户登录 ✅
func (s *userService) Login(req *types.UserLoginRequest) (string, error) {
	// 根据用户名获取用户信息
	user, err := s.repo.UserRepo.GetByUsername(req.Username)
	if err != nil {
		return "", fmt.Errorf("用户查询失败: %v", err)
	}
	if user == nil {
		return "", fmt.Errorf("用户 %s 不存在", req.Username)
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", fmt.Errorf("用户不存在或密码错误")
	}

	// 生成JWT token
	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", fmt.Errorf("生成Token失败: %v", err)
	}

	// 返回 token
	return token, nil
}

// GetUserByUsername 根据用户名获取用户 ✅
func (s *userService) GetUserByUsername(username string) (*auth.User, error) {
	return s.repo.UserRepo.GetByUsername(username)
}

// GetUserByID 根据ID获取用户  ✅
func (s *userService) GetUserByID(id uint) (*auth.User, error) {
	return s.repo.UserRepo.GetByID(id)
}

// CreateUser 创建用户 ✅
func (s *userService) CreateUser(user *auth.User) error {
	// 检查用户名是否已存在
	existingUser, err := s.repo.UserRepo.GetByUsername(user.Username)
	if err == nil && existingUser != nil {
		return fmt.Errorf("用户名 %s 已存在", user.Username)
	}

	// 创建用户
	return s.repo.UserRepo.Create(user)
}

// UpdateUser 更新用户 ✅
func (s *userService) UpdateUser(user *auth.User) error {
	return s.repo.UserRepo.Update(user)
}

// DeleteUser 删除用户 ✅
func (s *userService) DeleteUser(id uint) error {
	return s.repo.UserRepo.Delete(id)
}

// UpdateUserStatus 更新用户状态 ✅
func (s *userService) UpdateUserStatus(id uint, status int) error {
	return s.repo.UserRepo.UpdateStatus(id, status)
}

// QueryUsers 查询用户列表
func (s *userService) QueryUsers(query *types.UserQueryRequest) (*types.UserListResponse, error) {
	// 构建查询条件
	var conditions []interface{}
	whereClause := ""

	// 添加各种查询条件
	if query.Username != "" {
		whereClause += " AND user_name LIKE ?"
		conditions = append(conditions, "%"+query.Username+"%")
	}

	if query.Nickname != "" {
		whereClause += " AND nick_name LIKE ?"
		conditions = append(conditions, "%"+query.Nickname+"%")
	}

	if query.Email != "" {
		whereClause += " AND email LIKE ?"
		conditions = append(conditions, "%"+query.Email+"%")
	}

	if query.Phone != "" {
		whereClause += " AND phone LIKE ?"
		conditions = append(conditions, "%"+query.Phone+"%")
	}

	if query.Status != nil {
		whereClause += " AND status = ?"
		conditions = append(conditions, *query.Status)
	}

	if query.Sex != nil {
		whereClause += " AND sex = ?"
		conditions = append(conditions, *query.Sex)
	}

	// 移除开头的 AND
	if len(whereClause) > 0 {
		whereClause = whereClause[5:] // 去掉开头的 " AND"
	}

	// 计算总数
	var total int64
	queryBuilder := s.repo.UserRepo.Query()
	if whereClause != "" {
		// 添加条件到查询构建器
		queryBuilder = queryBuilder.Where(whereClause, conditions...)
	}

	err := queryBuilder.Count(&total).Error
	if err != nil {
		return nil, fmt.Errorf("查询用户总数失败: %v", err)
	}

	// 设置分页参数
	page := query.Page
	if page <= 0 {
		page = 1
	}

	pageSize := query.PageSize
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 构建排序语句
	orderBy := query.OrderBy
	if orderBy == "" {
		orderBy = "id"
	}

	sortOrder := query.SortOrder
	if sortOrder == "" {
		sortOrder = "desc"
	}

	// 查询用户列表
	var users []auth.User
	err = queryBuilder.Offset(offset).Limit(pageSize).Order(fmt.Sprintf("%s %s", orderBy, sortOrder)).Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("查询用户列表失败: %v", err)
	}

	// 返回结果
	return &types.UserListResponse{
		Total: total,
		Items: users,
		Page:  page,
		Size:  pageSize,
	}, nil
}

// QueryUsersWithCondition 查询用户
func (s *userService) QueryUsersWithCondition(query *types.UserQuery) (*types.PaginationResponse, error) {
	users, count, err := s.repo.UserRepo.QueryUsersWithCondition(query)
	if err != nil {
		return nil, err
	}

	// 计算总页数
	pages := int64(0)
	if query.PageSize > 0 {
		pages = (count + int64(query.PageSize) - 1) / int64(query.PageSize)
	}

	// 返回分页响应
	return &types.PaginationResponse{
		Total:    count,
		Page:     query.Page,
		PageSize: query.PageSize,
		Pages:    pages,
		Data:     users,
	}, nil
}

// BatchDeleteUser 批量删除用户(手动事务) ✅
func (s *userService) BatchDeleteUser(ids []uint) error {
	if len(ids) == 0 {
		return fmt.Errorf("用户ID 列表不能为空")
	}

	// 使用事务确保数据一致性
	tx := s.repo.UserRepo.BeginTransaction() // 开始事务
	if tx.Error != nil {
		return tx.Error // 如果开始事务失败，返回错误
	}

	// 执行批量删除操作
	if err := tx.Where("id IN (?)", ids).Delete(&auth.User{}).Error; err != nil {
		tx.Rollback() // 如果删除失败，回滚事务
		return err
	}

	return tx.Commit().Error // 提交事务或返回错误
}

// BatchDeleteUser2 批量删除用户(简化事务) ✅
func (s *userService) BatchDeleteUser2(ids []uint) error {
	if len(ids) == 0 {
		return fmt.Errorf("用户ID 列表不能为空")
	}

	// 使用 gorm 的 Transaction 方法
	return s.repo.UserRepo.Transaction(func(tx *gorm.DB) error {
		// 执行批量删除操作
		if err := tx.Where("id IN (?)", ids).Delete(&auth.User{}).Error; err != nil {
			return fmt.Errorf("批量删除用户失败: %v", err)
		}
		return nil
	})
}
