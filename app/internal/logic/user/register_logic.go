package userlogic

import (
	"context"

	"go-zero-box-rpc/app/internal/svc"
	"go-zero-box-rpc/app/rpc/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user_rpc.UserRegisterReq) (*user_rpc.UserRegisterResp, error) {
	// todo: add your logic here and delete this line

	return &user_rpc.UserRegisterResp{}, nil
}
