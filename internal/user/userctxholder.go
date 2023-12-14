package common

// TODO 添加接口抽象方法
type UserCtxHolder interface {
}

type UserCtx struct {
	Id       int64
	Username string
}
