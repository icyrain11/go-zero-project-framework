package error

const (
	RedisGetErrorCode = 50001 + iota
	RedisExpireErrorCode
	RedisDelErrorCode
	JsonUnmarshalErrorCode
	DocumentFromReaderErrorCode
)

const (
	ServerInternalErrorMsg = "系统繁忙"
)
