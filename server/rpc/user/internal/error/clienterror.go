package error

const (
	BadRequestCode = 40000 + iota
	InvalidCode
)

const (
	BadRequest = "参数错误"
)

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
