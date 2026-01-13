package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logc"
	"go-zero-box-rpc/app/internal/command"
	"go-zero-box-rpc/app/internal/queue"
	"go-zero-box-rpc/pkg/asynqx"
	"log"

	"go-zero-box-rpc/app/internal/config"
	helloServer "go-zero-box-rpc/app/internal/server/hello"
	userServer "go-zero-box-rpc/app/internal/server/user"
	"go-zero-box-rpc/app/rpc/user_rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	configFile = flag.String("conf", "app/etc/app.yaml", "the config file")

	rootCmd = &cobra.Command{
		Use:                "app",
		DisableFlagParsing: true,
		Hidden:             true,
	}
)

func main() {
	flag.Parse()

	var c *config.Config
	conf.MustLoad(*configFile, &c)

	logc.MustSetup(c.Server.Log)
	app := initApp(c)
	rootCmd.AddCommand(serverRpc(app), serverQueue(app), serverScheduler(app), serverAll(app))
	rootCmd.AddCommand(command.RegisterHandlerScript(app.command)...)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("execute core service failed, %s\n", err.Error())
	}
}

func serverRpc(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:rpc",
		Short: "启动rpc服务",
		Run: func(cmd *cobra.Command, args []string) {
			server := zrpc.MustNewServer(app.config.Server, func(grpcServer *grpc.Server) {
				user_rpc.RegisterHelloServer(grpcServer, helloServer.NewHelloServer(app.svcCtx))
				user_rpc.RegisterUserServer(grpcServer, userServer.NewUserServer(app.svcCtx))

				if app.config.Server.Mode == service.DevMode || app.config.Server.Mode == service.TestMode {
					reflection.Register(grpcServer)
				}
			})
			defer server.Stop()
			log.Printf("[server:rpc] start success at %s...\n", app.config.Server.ListenOn)
			server.Start()
		},
	}
}

func serverQueue(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:queue",
		Short: "启动队列服务",
		Run: func(cmd *cobra.Command, args []string) {
			serviceGroup := service.NewServiceGroup()
			defer serviceGroup.Stop()

			handlers := queue.RegisterHandlerQueue(app.queue)
			for _, v := range handlers {
				serviceGroup.Add(asynqx.NewQueue(app.config.Asynqx, v))
			}
			serviceGroup.Start()
			select {}
		},
	}
}

func serverScheduler(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:scheduler",
		Short: "启动计划任务服务",
		Run: func(cmd *cobra.Command, args []string) {
			serviceGroup := service.NewServiceGroup()
			defer serviceGroup.Stop()

			handlers := command.RegisterHandlerScheduler(app.command)
			serviceGroup.Add(asynqx.NewScheduler(app.config.Asynqx, handlers))
			for _, v := range handlers {
				serviceGroup.Add(asynqx.NewQueue(app.config.Asynqx, v))
			}

			serviceGroup.Start()
		},
	}
}

func serverAll(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:all",
		Short: "单体式服务，包含api、队列、计划任务",
		Run: func(cmd *cobra.Command, args []string) {
			go serverRpc(app).Run(serverRpc(app), []string{})
			go serverQueue(app).Run(serverQueue(app), []string{})
			go serverScheduler(app).Run(serverScheduler(app), []string{})
			select {}
		},
	}
}
