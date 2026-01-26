package constant

// 错误码
const (
	ErrCodeSuccess        = 200
	ErrCodeInvalidParam   = 400
	ErrCodeUserExist      = 401
	ErrCodeUserNotFound   = 402
	ErrCodePasswordError  = 403
	ErrCodeTokenInvalid   = 404
	ErrCodeServerInternal = 500
)

// 错误信息
const (
	MsgSuccess        = "success"
	MsgInvalidParam   = "参数错误"
	MsgUserExist      = "用户已存在"
	MsgUserNotFound   = "用户不存在"
	MsgPasswordError  = "密码错误"
	MsgTokenInvalid   = "token无效"
	MsgServerInternal = "服务器内部错误"
)

// Redis Key前缀（单Token简化）
const (
	RedisKeyToken = "token:"       // 单Token存储前缀：key=token:xxx, value=userID
	RedisKeyBlack = "token:black:" // Token黑名单前缀（退出后防复用）
	// 过期时间也可以直接使用jwt配置的过期时间
	RedisTokenExpire = 2 * 60 * 60 // Redis中token 的过期时间2小时，单位：秒

)
