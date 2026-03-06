package hellologic

import (
	"context"

	"github.com/prf16/go-zero-box-rpc/app/api/user"

	"github.com/prf16/go-zero-box-rpc/app/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorldLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWorldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorldLogic {
	return &WorldLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WorldLogic) World(in *user.HelloWorldReq) (*user.HelloWorldResp, error) {
	// todo: add your logic here and delete this line

	return &user.HelloWorldResp{
		Message: "hello world",
	}, nil
}
