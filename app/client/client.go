package main

import (
	"context"
	"fmt"
	"go-zero-box-rpc/app/rpc/user_rpc"

	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	ctx := context.Background()
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "127.0.0.1:8080",
	}).Conn()

	userInfoResp, err := user_rpc.NewHelloClient(client).World(ctx, &user_rpc.HelloWorldReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfoResp)
}
