package config

import (
	"github.com/google/wire"
	"github.com/prf16/go-zero-box-rpc/pkg/asynqx"
	"github.com/prf16/go-zero-box-rpc/pkg/database"
	"github.com/prf16/go-zero-box-rpc/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
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
