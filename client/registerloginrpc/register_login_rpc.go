// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package registerloginrpc

import (
	"context"

	"github.com/go-bineanshi/usercenter-rpc/types/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseMessage       = usercenter.BaseMessage
	ChangePasswordReq = usercenter.ChangePasswordReq
	LoginUserReq      = usercenter.LoginUserReq
	RegisterUserReq   = usercenter.RegisterUserReq
	UserInfo          = usercenter.UserInfo

	RegisterLoginRpc interface {
		RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*UserInfo, error)
		LoginUser(ctx context.Context, in *LoginUserReq, opts ...grpc.CallOption) (*UserInfo, error)
		ChangeUserPassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*BaseMessage, error)
	}

	defaultRegisterLoginRpc struct {
		cli zrpc.Client
	}
)

func NewRegisterLoginRpc(cli zrpc.Client) RegisterLoginRpc {
	return &defaultRegisterLoginRpc{
		cli: cli,
	}
}

func (m *defaultRegisterLoginRpc) RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*UserInfo, error) {
	client := usercenter.NewRegisterLoginRpcClient(m.cli.Conn())
	return client.RegisterUser(ctx, in, opts...)
}

func (m *defaultRegisterLoginRpc) LoginUser(ctx context.Context, in *LoginUserReq, opts ...grpc.CallOption) (*UserInfo, error) {
	client := usercenter.NewRegisterLoginRpcClient(m.cli.Conn())
	return client.LoginUser(ctx, in, opts...)
}

func (m *defaultRegisterLoginRpc) ChangeUserPassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*BaseMessage, error) {
	client := usercenter.NewRegisterLoginRpcClient(m.cli.Conn())
	return client.ChangeUserPassword(ctx, in, opts...)
}
