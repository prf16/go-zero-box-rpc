package main

import (
	"context"
	"fmt"
	"github.com/prf16/go-zero-box-rpc/api/user"

	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	ctx := context.Background()
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "127.0.0.1:8080",
	}).Conn()

	userInfoResp, err := user.NewHelloClient(client).World(ctx, &user.HelloWorldReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfoResp)
}
