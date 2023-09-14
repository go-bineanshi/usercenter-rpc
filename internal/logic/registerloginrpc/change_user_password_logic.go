package registerloginrpclogic

import (
  "context"
  "fmt"
  "github.com/go-bineanshi/usercenter-rpc/utils"

  "github.com/go-bineanshi/usercenter-rpc/internal/svc"
  "github.com/go-bineanshi/usercenter-rpc/types/usercenter"

  "github.com/zeromicro/go-zero/core/logx"
)

type ChangeUserPasswordLogic struct {
  ctx    context.Context
  svcCtx *svc.ServiceContext
  logx.Logger
}

func NewChangeUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeUserPasswordLogic {
  return &ChangeUserPasswordLogic{
    ctx:    ctx,
    svcCtx: svcCtx,
    Logger: logx.WithContext(ctx),
  }
}

func (l *ChangeUserPasswordLogic) ChangeUserPassword(in *usercenter.ChangePasswordReq) (*usercenter.BaseMessage, error) {
  user, err := l.svcCtx.DB.User.Get(l.ctx, in.Id)
  if err != nil {
    return nil, fmt.Errorf("用户不存在")
  }
  if ok := utils.CheckPassword(in.Oldpassword, user.Password); !ok {
    return nil, fmt.Errorf("密码不正确")
  }
  password, err := utils.GeneratePassword(in.Password)
  if err != nil {
    return nil, fmt.Errorf("加密失败")
  }
  _, err = user.Update().
    SetPassword(password).
    Save(l.ctx)
  if err != nil {
    return nil, fmt.Errorf("密码更新失败")
  }
  return &usercenter.BaseMessage{}, nil
}
