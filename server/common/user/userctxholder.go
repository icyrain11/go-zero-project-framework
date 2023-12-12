package common

type UserCtx struct {
	Id       int64
	Username string
}

func GetUserCtx(token string) (userCtx *UserCtx) {

	//从Redis中获取

	return
}
