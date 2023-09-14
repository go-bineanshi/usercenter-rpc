package registerloginrpclogic

import (
  "context"
  "fmt"
  "github.com/go-bineanshi/usercenter-rpc/internal/svc"
  "github.com/go-bineanshi/usercenter-rpc/types/usercenter"
  "github.com/go-bineanshi/usercenter-rpc/utils"
  "github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
  ctx    context.Context
  svcCtx *svc.ServiceContext
  logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
  return &RegisterUserLogic{
    ctx:    ctx,
    svcCtx: svcCtx,
    Logger: logx.WithContext(ctx),
  }
}

func (l *RegisterUserLogic) RegisterUser(in *usercenter.RegisterUserReq) (*usercenter.UserInfo, error) {
  // 开启事务
  tx, err := l.svcCtx.DB.Tx(l.ctx)

  if err != nil {
    return nil, fmt.Errorf("failed starting transaction: %v", err)
  }

  defer tx.Rollback() // 确保在错误发生时回滚事务

  // 创建用户名称类型的 用户授权
  userAuth, err := tx.UserAuth.Create().
    SetAccount(in.Username).
    SetProvider(utils.CheckLoginnameType(in.Username)).
    Save(l.ctx)
  if err != nil {
    return nil, fmt.Errorf("failed creating user authentication: %v", err)
  }

  password, err := utils.GeneratePassword(in.Password)
  if err != nil {
    return nil, fmt.Errorf("加密失败: %v", err)
  }

  // 创建用户
  user, err := tx.User.Create().
    SetNickname(in.Username).
    SetPassword(password).
    AddAuths(userAuth).
    Save(l.ctx)
  if err != nil {
    return nil, fmt.Errorf("failed creating user: %w", err)
  }
  // 事务提交
  if err := tx.Commit(); err != nil {
    return nil, fmt.Errorf("failed committing transaction: %v", err)
  }

  createdAt := user.CreatedAt.UnixNano()
  updatedAt := user.UpdatedAt.UnixNano()
  return &usercenter.UserInfo{
    Id:        &user.ID,
    Nickname:  &user.Nickname,
    Avater:    &user.Avater,
    CreatedAt: &createdAt,
    UpdatedAt: &updatedAt,
  }, nil
}
