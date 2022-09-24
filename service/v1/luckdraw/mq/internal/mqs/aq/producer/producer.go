package producer

import (
	"context"
	"luck_draw-server/common/asynqmq"

	"github.com/hibiken/asynq"
)

func NewAsynqProducer(ctx context.Context, redisClientOpt *asynq.RedisClientOpt) *asynqmq.AsynqProducer {
	return asynqmq.NewAsynqProducer(ctx, redisClientOpt)
}
