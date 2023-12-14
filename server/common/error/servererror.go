package error

const (
	RedisGetErrorCode = 50001 + iota
	RedisExpireErrorCode
	RedisDelErrorCode
	JsonUnmarshalErrorCode
)

const (
	ServerInternalErrorMsg = "系统繁忙"
)
