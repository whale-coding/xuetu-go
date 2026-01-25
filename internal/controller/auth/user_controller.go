package auth

import (
	"log"
	"net/http"
	"strconv"
	"xuetu-project/internal/model/auth"
	"xuetu-project/internal/pkg/response"
	"xuetu-project/internal/service"
	"xuetu-project/internal/types"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	GetUserByUsername(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUserStatus(ctx *gin.Context)
	QueryUsers(ctx *gin.Context)
	QueryUsersWithCondition(ctx *gin.Context)
	BatchDeleteUsers(ctx *gin.Context)
}

type userController struct {
	svc *service.Service
}

func NewUserController(svc *service.Service) UserController {
	return &userController{svc: svc}
}

// Register 用户注册
func (c *userController) Register(ctx *gin.Context) {
	// 获取请求参数
	var req types.UserRegisterRequest

	// 参数绑定与验证
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "参数验证失败: " + err.Error()})
		return
	}

	log.Printf("注册请求参数: %+v\n", req)

	// 调用服务层的注册逻辑
	if err := c.svc.UserService.Register(&req); err != nil {
		ctx.JSON(500, gin.H{"error": "注册失败: " + err.Error()})
		return
	}
	// 返回成功响应,统一返回格式
	response.Success(ctx, "注册成功")
}

// Login 用户登录
func (c *userController) Login(ctx *gin.Context) {
	// 获取请求参数
	var req types.UserLoginRequest

	// 参数绑定与验证
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "参数验证失败: " + err.Error()})
		return
	}
	log.Printf("登录请求参数: %+v\n", req)

	// 调用服务层的登录逻辑
	token, err := c.svc.UserService.Login(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "登录失败: " + err.Error()})
		return
	}
	// 将 string 格式的token转换为LoginResponse格式
	res := types.LoginResponse{
		Token: token,
	}

	response.Success(ctx, res)
}

// CreateUser 创建用户
func (c *userController) CreateUser(ctx *gin.Context) {
	// 获取请求参数
	var req types.UserCreateRequest
	// 参数绑定与验证
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "参数验证失败: " + err.Error()})
		return
	}
	log.Printf("创建用户请求参数: %+v\n", req)

	// 将请求参数转换为 User 模型
	user := &auth.User{
		Username:  req.Username,
		Nickname:  req.Nickname,
		Email:     req.Email,
		Phone:     req.Phone,
		Sex:       req.Sex,
		Avatar:    req.Avatar,
		Introduce: req.Introduce,
		ExtJson:   req.ExtJson,
		Status:    1, // 默认状态为正常
	}

	// 调用服务层创建用户
	if err := c.svc.UserService.CreateUser(user); err != nil {
		ctx.JSON(500, gin.H{"error": "用户创建失败: " + err.Error()})
		return
	}

	response.Success(ctx, "用户创建成功")
}

// GetUserByUsername 根据用户名获取用户
// 请求示例: GET /api/user?username=exampleUser
func (c *userController) GetUserByUsername(ctx *gin.Context) {
	// 从查询参数中获取用户名
	username := ctx.Query("username")
	log.Printf("查询参数: %+v\n", username)

	if username == "" {
		ctx.JSON(400, gin.H{"error": "用户名不能为空: "})
		return
	}

	// 调用服务层获取用户信息
	user, err := c.svc.UserService.GetUserByUsername(username)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "用户信息获取失败: " + err.Error()})
		return
	}
	if user == nil {
		ctx.JSON(404, gin.H{"error": "用户不存在"})
		return
	}

	response.Success(ctx, user)
}

// GetUserByID 根据ID获取用户
// 请求示例: GET /api/user/1919728947180916742
func (c *userController) GetUserByID(ctx *gin.Context) {
	// 从路径参数中获取用户 ID
	userIDStr := ctx.Param("id")
	log.Printf("路径参数: %+v\n", userIDStr)

	// 将字符串转换为整数
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户 ID"})
		return
	}

	// 调用服务层获取用户信息
	user, err := c.svc.UserService.GetUserByID(uint(userID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "用户信息获取失败: " + err.Error()})
		return
	}
	if user == nil {
		ctx.JSON(404, gin.H{"error": "用户不存在"})
		return
	}

	response.Success(ctx, user)
}

// UpdateUser 更新用户
// 请求示例: PUT /api/user/1919728947180916742
func (c *userController) UpdateUser(ctx *gin.Context) {
	// 从路径参数中获取用户 ID
	idStr := ctx.Param("id")
	log.Printf("路径参数: %+v\n", idStr)

	// 将字符串转换为整数
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户 ID"})
		return
	}

	// 获取请求参数
	var req types.UserUpdateRequest

	// 参数绑定与验证
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "参数验证失败: " + err.Error()})
		return
	}

	// 获取现有用户信息
	existingUser, err := c.svc.UserService.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "获取用户失败: " + err.Error()})
		return
	}
	log.Printf("现有用户信息: %+v\n", existingUser)

	// 更新用户信息
	existingUser.Username = req.Username
	existingUser.Nickname = req.Nickname
	existingUser.Email = req.Email
	existingUser.Phone = req.Phone
	// 注意：实际项目中需要对密码进行哈希处理
	if req.Password != "" {
		// 对密码进行哈希处理
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "密码哈希处理失败: " + err.Error()})
		}
		existingUser.Password = string(hashedPassword)
	}
	existingUser.Sex = req.Sex
	existingUser.Avatar = req.Avatar
	existingUser.Introduce = req.Introduce
	existingUser.ExtJson = req.ExtJson
	if req.Status != 0 {
		existingUser.Status = req.Status
	}

	// 调用服务层更新用户信息
	if err := c.svc.UserService.UpdateUser(existingUser); err != nil {
		ctx.JSON(400, gin.H{"error": "更新用户失败: " + err.Error()})
		return
	}

	response.Success(ctx, "用户更新成功")
}

// DeleteUser 删除用户
// 请求示例: DELETE /api/user/1919728947180916742
func (c *userController) DeleteUser(ctx *gin.Context) {
	// 从路径参数中获取用户 ID
	idStr := ctx.Param("id")
	log.Printf("路径参数: %+v\n", idStr)

	// 将字符串转换为整数
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户 ID"})
		return
	}

	// 调用服务层删除用户
	if err := c.svc.UserService.DeleteUser(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": "删除用户失败: " + err.Error()})
		return
	}

	response.Success(ctx, "用户删除成功")
}

// UpdateUserStatus 更新用户状态
// 请求示例: PATCH /api/user/1919728947180916742/status
func (c *userController) UpdateUserStatus(ctx *gin.Context) {
	// 从路径参数中获取用户 ID
	idStr := ctx.Param("id")
	log.Printf("路径参数: %+v\n", idStr)

	// 将字符串转换为整数
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户 ID"})
		return
	}

	// 获取请求参数
	var req struct {
		Status int `json:"status" binding:"oneof=0 1"`
	}

	// 参数绑定与验证
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数验证失败: " + err.Error()})
		return
	}
	log.Printf("请求参数: %+v\n", req)

	// 验证状态值
	if req.Status != 0 && req.Status != 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "状态值无效"})
		return
	}

	// 调用服务层更新用户状态
	if err := c.svc.UserService.UpdateUserStatus(uint(id), req.Status); err != nil {
		ctx.JSON(500, gin.H{"error": "更新用户状态失败: " + err.Error()})
		return
	}

	response.Success(ctx, "用户状态更新成功")
}

// QueryUsers 查询用户列表(Get请求方式)
// 请求示例: GET /api/user?page=1&page_size=10&username=admin&status=1
func (c *userController) QueryUsers(ctx *gin.Context) {
	// 绑定查询参数
	var query types.UserQueryRequest

	// 参数绑定与验证
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数验证失败: " + err.Error()})
		return
	}

	log.Printf("查询参数: %+v\n", query)

	// 调用服务层查询用户
	result, err := c.svc.UserService.QueryUsers(&query)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "查询用户失败: " + err.Error()})
		return
	}

	response.Success(ctx, result)
}

// BatchDeleteUsers 批量删除用户
// 请求示例: POST /api/user/batch-delete
func (c *userController) BatchDeleteUsers(ctx *gin.Context) {
	// 获取请求参数，即从请求体中获取用户ID列表
	var req types.BatchDeleteRequest

	// 参数绑定与验证
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数验证失败: " + err.Error()})
		return
	}
	log.Printf("批量删除参数: %+v\n", req)

	// 调用服务层批量删除用户
	if err := c.svc.UserService.BatchDeleteUser2(req.IDs); err != nil {
		ctx.JSON(500, gin.H{"error": "批量删除用户失败: " + err.Error()})
		return
	}

	response.Success(ctx, "用户批量删除成功")
}

// QueryUsersWithCondition 分页查询用户列表（POST请求方式）
func (c *userController) QueryUsersWithCondition(ctx *gin.Context) {
	// 获取请求参数
	var query types.UserQuery

	// 参数绑定与验证
	if err := ctx.ShouldBindJSON(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数验证失败: " + err.Error()})
		return
	}

	// 设置默认值
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 10 // 默认每页10条
	}
	log.Printf("查询参数: %+v\n", query)

	// 调用服务层查询用户
	result, err := c.svc.UserService.QueryUsersWithCondition(&query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询用户失败: " + err.Error()})
		return
	}

	response.Success(ctx, result)
	//ctx.JSON(http.StatusOK, gin.H{
	//	"data":      users,
	//	"total":     count,
	//	"page":      query.Page,
	//	"page_size": query.PageSize,
	//})
}
