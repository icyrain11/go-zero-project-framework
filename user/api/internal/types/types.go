// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type GetUserInfoReq struct {
	Id int64 `path:"id"`
}

type GetUserInfoResp struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}