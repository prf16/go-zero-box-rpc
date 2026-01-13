package config

import (
	"github.com/google/wire"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-box-rpc/pkg/asynqx"
	"go-zero-box-rpc/pkg/database"
	"go-zero-box-rpc/pkg/redis"
)

var Provider = wire.NewSet(
	wire.FieldsOf(new(*Config), "Database", "Redis", "Asynqx"),
)

type Config struct {
	Server   zrpc.RpcServerConf
	Database *database.Config
	Redis    *redis.Config
	Asynqx   *asynqx.Config
}
