package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-box-rpc/app/rpc"
)

func main() {
	ctx := context.Background()
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "127.0.0.1:8080",
	})

	userInfoResp, err := rpc.NewHelloClient(conn.Conn()).World(ctx, &rpc.HelloWorldReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfoResp)
}
