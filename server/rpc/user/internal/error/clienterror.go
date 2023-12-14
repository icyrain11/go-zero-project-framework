package error

const (
	UserNotFoundErrorCode = 40100 + iota
	UserPasswordErrorCode
	UserNotLoginErrorCode
)

const (
	UserNotFoundErrorMsg = "用户不存在"
	UserPasswordErrorMsg = "密码错误"
	UserNotLoginErrorMsg = "用户未登录"
)
