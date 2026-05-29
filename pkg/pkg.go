package pkg

import (
	"github.com/prf16/go-zero-box-rpc/pkg/asynqx"
	"github.com/prf16/go-zero-box-rpc/pkg/database"
	"github.com/prf16/go-zero-box-rpc/pkg/redis"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewPkg,
	database.Provider,
	redis.Provider,
	asynqx.Provider,
)

type Pkg struct {
	Database *database.Database
	Redis    *redis.Redis
	Asynqx   *asynqx.Asynq
}

func NewPkg(database *database.Database, redis *redis.Redis, asynqx *asynqx.Asynq) *Pkg {
	return &Pkg{Database: database, Redis: redis, Asynqx: asynqx}
}
