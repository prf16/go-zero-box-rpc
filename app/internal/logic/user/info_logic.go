package userlogic

import (
	"context"

	"go-zero-box-rpc/app/internal/svc"
	"go-zero-box-rpc/app/rpc/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *user_rpc.UserInfoReq) (*user_rpc.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user_rpc.UserInfoResp{
		User: &user_rpc.UserModel{
			ID:       1,
			NickName: "prf16",
		},
	}, nil
}
