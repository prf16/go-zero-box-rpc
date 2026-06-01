package app

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/prf16/go-zero-box-rpc/api/user"
	"github.com/prf16/go-zero-box-rpc/app/internal/svc/queue"
	"github.com/prf16/go-zero-box-rpc/pkg/asynqx"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/prf16/go-zero-box-rpc/app/internal/config"
	helloServer "github.com/prf16/go-zero-box-rpc/app/internal/server/hello"
	userServer "github.com/prf16/go-zero-box-rpc/app/internal/server/user"

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

func Start() {
	flag.Parse()

	var c *config.Config
	err := conf.Load(*configFile, &c)
	if err != nil {
		err = conf.Load("etc/app.yaml", &c)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
	}

	logc.MustSetup(c.Server.Log)
	defer logc.Close()

	app := initApp(c)

	rootCmd.AddCommand(
		serverRpc(app),
		serverQueue(app),
		serverScheduler(app),
		serverAll(app),
	)

	for _, v := range app.svcCtx.Command.Register() {
		rootCmd.AddCommand(v.Command)
	}

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
				user.RegisterHelloServer(grpcServer, helloServer.NewHelloServer(app.svcCtx))
				user.RegisterUserServer(grpcServer, userServer.NewUserServer(app.svcCtx))

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

			handlers := queue.Register(app.svcCtx.Queue)
			for _, v := range handlers {
				serviceGroup.Add(asynqx.NewQueue(app.config.Redis, v))
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

			var handlers []*asynqx.Handler
			for _, v := range app.svcCtx.Command.Register() {
				if v.Scheduler == "" {
					continue
				}

				handlers = append(handlers, &asynqx.Handler{
					Type:      v.Command.Use,
					Scheduler: v.Scheduler,
					Handler: func(ctx context.Context, task *asynq.Task) error {
						v.Command.Run(v.Command, nil)
						return nil
					},
				})
			}

			serviceGroup.Add(asynqx.NewScheduler(app.config.Redis, handlers))
			for _, v := range handlers {
				serviceGroup.Add(asynqx.NewQueue(app.config.Redis, v))
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
