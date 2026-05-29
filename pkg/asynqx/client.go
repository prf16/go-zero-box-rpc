package asynqx

import (
	"github.com/prf16/go-zero-box-rpc/pkg/redis"

	"github.com/hibiken/asynq"
)

func NewClient(c *redis.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     c.Host,
		Password: c.Pass,
	})
}

func NewTask(typename string, payload []byte, opts ...asynq.Option) *asynq.Task {
	return asynq.NewTask(typename, payload, opts...)
}
