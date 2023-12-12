// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"
	user2 "go-zero-demo/server/rpc/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckLoginStatusReq  = user2.CheckLoginStatusReq
	CheckLoginStatusResp = user2.CheckLoginStatusResp
	LoginReq             = user2.LoginReq
	LoginResp            = user2.LoginResp

	User interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		CheckLoginStatus(ctx context.Context, in *CheckLoginStatusReq, opts ...grpc.CallOption) (*CheckLoginStatusResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user2.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) CheckLoginStatus(ctx context.Context, in *CheckLoginStatusReq, opts ...grpc.CallOption) (*CheckLoginStatusResp, error) {
	client := user2.NewUserClient(m.cli.Conn())
	return client.CheckLoginStatus(ctx, in, opts...)
}
