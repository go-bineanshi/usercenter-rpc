package registerloginrpclogic

import (
  "context"
  "fmt"
  "github.com/go-bineanshi/usercenter-rpc/ent/userauth"
  "github.com/go-bineanshi/usercenter-rpc/internal/svc"
  "github.com/go-bineanshi/usercenter-rpc/types/usercenter"
  "github.com/go-bineanshi/usercenter-rpc/utils"

  "github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
  ctx    context.Context
  svcCtx *svc.ServiceContext
  logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
  return &LoginUserLogic{
    ctx:    ctx,
    svcCtx: svcCtx,
    Logger: logx.WithContext(ctx),
  }
}

func (l *LoginUserLogic) LoginUser(in *usercenter.LoginUserReq) (*usercenter.UserInfo, error) {
  auth, err := l.svcCtx.DB.UserAuth.Query().
    Where(userauth.AccountEQ(in.Loginname), userauth.ProviderEQ(utils.CheckLoginnameType(in.Loginname))).
    WithUser().Only(l.ctx)

  if err != nil {
    return nil, fmt.Errorf("user not found", nil)
  }

  if auth.Status != 1 {
    return nil, fmt.Errorf("该认证方式不可用")
  }
  user := auth.Edges.User

  if ok := utils.CheckPassword(in.Password, user.Password); !ok {
    return nil, fmt.Errorf("密码不正确")
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
