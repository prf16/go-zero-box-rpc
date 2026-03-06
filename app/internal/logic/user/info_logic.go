package userlogic

import (
	"context"

	"github.com/prf16/go-zero-box-rpc/app/api/user"

	"github.com/prf16/go-zero-box-rpc/app/internal/svc"

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

func (l *InfoLogic) Info(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoResp{
		User: &user.UserModel{
			ID:       1,
			NickName: "prf16",
		},
	}, nil
}
