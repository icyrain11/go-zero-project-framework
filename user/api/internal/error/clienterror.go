package error

const (
	BadRequestCode = 40000 + iota
	InvalidCode
)

const (
	BadRequestMsg = "参数错误"
)

const (
	UserNotFoundCode = 40100 + iota
	UserPasswordErrorCode
)

const (
	UserNotFound      = "用户不存在"
	UserPasswordError = "密码错误"
)
