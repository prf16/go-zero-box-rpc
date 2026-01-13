package hellologic

import (
	"context"

	"go-zero-box-rpc/app/internal/svc"
	"go-zero-box-rpc/app/rpc"

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

func (l *WorldLogic) World(in *rpc.HelloWorldReq) (*rpc.HelloWorldResp, error) {
	// todo: add your logic here and delete this line

	return &rpc.HelloWorldResp{
		Message: "Hello World",
	}, nil
}
