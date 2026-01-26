package types

import (
	"xuetu-project/internal/model"
)

// UserRegisterRequest 注册请求
type UserRegisterRequest struct {
	Username  string `json:"username" binding:"required,min=3,max=20"`  // 用户名,必须，长度3-20
	Password  string `json:"password" binding:"required,min=6,max=20"`  // 密码,必须，长度6-20
	Nickname  string `json:"nickname" binding:"omitempty,min=2,max=20"` // 昵称,可选，长度2-20
	Email     string `json:"email" binding:"omitempty,email"`           // 邮箱,可选，必须是有效邮箱格式
	Phone     string `json:"phone" binding:"omitempty"`                 // 电话,可选
	Sex       int    `json:"sex" binding:"omitempty,oneof=0 1 2"`       // 性别,可选，0-未知,1-男,2-女
	Avatar    string `json:"avatar" binding:"omitempty"`                // 头像,可选
	Introduce string `json:"introduce" binding:"omitempty,max=255"`     // 介绍,可选，最大长度255
	ExtJson   string `json:"extJson" binding:"omitempty"`               // 扩展信息,可选
}

// UserLoginRequest 登录请求
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 定义登录成功后的返回格式
type LoginResponse struct {
	Token string `json:"token"`
}

// UserUpdateInfoRequest 更新用户信息请求
type UserUpdateInfoRequest struct {
	Nickname  string `json:"nickname" binding:"omitempty,min=2,max=20"` // 昵称,可选，长度2-20
	Email     string `json:"email" binding:"omitempty,email"`           // 邮箱,可选，必须是有效邮箱格式
	Phone     string `json:"phone" binding:"omitempty"`                 // 电话,可选
	Sex       int    `json:"sex" binding:"omitempty,oneof=0 1"`         // 性别,可选，0-男,1-女
	Avatar    string `json:"avatar" binding:"omitempty"`                // 头像,可选
	Introduce string `json:"introduce" binding:"omitempty,max=255"`     // 介绍,可选，最大长度255
}

// UserInfoResponse 用户信息响应
type UserInfoResponse struct {
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Sex       int    `json:"sex"`
	Avatar    string `json:"avatar"`
	Introduce string `json:"introduce"`
}

//// Service DTO - 服务层内部使用的数据结构
//type UserDTO struct {
//	ID       uint   `json:"id"`
//	Username string `json:"username"`
//	Nickname string `json:"nickname"`
//	Email    string `json:"email"`
//}

//// 将Model转换为DTO
//func ModelToDTO(user *auth.User) *UserDTO {
//	return &UserDTO{
//		ID:       user.ID,
//		Username: user.Username,
//		Nickname: user.Nickname,
//		Email:    user.Email,
//	}
//}
//

// UserCreateRequest 用户创建请求
type UserCreateRequest struct {
	Username  string `json:"userName" binding:"required,min=3,max=50"`
	Nickname  string `json:"nickName" binding:"required,min=2,max=50"`
	Email     string `json:"email" binding:"omitempty,email"`
	Phone     string `json:"phone" binding:"omitempty"`
	Password  string `json:"password" binding:"required,min=6,max=100"`
	Sex       int    `json:"sex" binding:"omitempty,oneof=0 1 2"`
	Avatar    string `json:"avatar" binding:"omitempty,url"`
	Introduce string `json:"introduce" binding:"omitempty,max=255"`
	ExtJson   string `json:"extJson" binding:"omitempty"`
}

// UserUpdateRequest 用户更新请求
type UserUpdateRequest struct {
	Username  string `json:"userName" binding:"omitempty,min=3,max=50"`
	Nickname  string `json:"nickName" binding:"omitempty,min=2,max=50"`
	Email     string `json:"email" binding:"omitempty,email"`
	Phone     string `json:"phone" binding:"omitempty"`
	Password  string `json:"password" binding:"omitempty,min=6,max=100"`
	Sex       int    `json:"sex" binding:"omitempty,oneof=0 1 2"`
	Avatar    string `json:"avatar" binding:"omitempty,url"`
	Introduce string `json:"introduce" binding:"omitempty,max=255"`
	ExtJson   string `json:"extJson" binding:"omitempty"`
	Status    int    `json:"status" binding:"omitempty,oneof=0 1"`
}

// UserQueryRequest 用户查询请求
type UserQueryRequest struct {
	Username string `form:"username"` // 用户名
	Nickname string `form:"nickname"` // 昵称
	Email    string `form:"email"`    // 邮箱
	Phone    string `form:"phone"`    // 电话
	Status   *int   `form:"status"`   // 状态
	Sex      *int   `form:"sex"`      // 性别
	//Page      int    `form:"page,default=1"`       // 页码
	//PageSize  int    `form:"page_size,default=10"` // 每页大小
	OrderBy           string `form:"order_by"`   // 排序字段
	SortOrder         string `form:"sort_order"` // 排序方向: asc/desc
	PaginationRequest        // 分页参数：页码和每页大小
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	Total int64        `json:"total"`
	Items []model.User `json:"items"`
	Page  int          `json:"page"`
	Size  int          `json:"size"`
}

// BatchDeleteRequest 批量删除请求
type BatchDeleteRequest struct {
	IDs []uint `json:"ids" binding:"required,min=1"`
}

type UserQuery struct {
	Username string `json:"username,omitempty"`
	Status   int    `json:"status,omitempty"`
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
}
