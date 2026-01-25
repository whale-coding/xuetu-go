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
