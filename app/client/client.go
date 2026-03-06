package main

import (
	"context"
	"fmt"

	user2 "github.com/prf16/go-zero-box-rpc/app/api/user"

	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	ctx := context.Background()
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "127.0.0.1:8080",
	}).Conn()

	userInfoResp, err := user2.NewHelloClient(client).World(ctx, &user2.HelloWorldReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfoResp)
}
