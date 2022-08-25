package notification

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type RedisSubscribableClient interface {
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
}

type RedisSubsriber struct {
	client RedisSubscribableClient
}

func NewRedisSubscriber(client RedisSubscribableClient) *RedisSubsriber {
	return &RedisSubsriber{
		client: client,
	}
}
