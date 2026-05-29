package main

import (
	"flag"
	"fmt"

	"github.com/prf16/go-zero-box-rpc/api/user"
	"github.com/prf16/go-zero-box-rpc/app/internal/config"
	helloServer "github.com/prf16/go-zero-box-rpc/app/internal/server/hello"
	userServer "github.com/prf16/go-zero-box-rpc/app/internal/server/user"
	"github.com/prf16/go-zero-box-rpc/app/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterHelloServer(grpcServer, helloServer.NewHelloServer(ctx))
		user.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
