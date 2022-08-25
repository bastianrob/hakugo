package notification

import (
	"context"

	"github.com/go-redis/redis/v9"
)

func (rs *RedisSubsriber) CustomerRegistrationStarted() <-chan *redis.Message {
	return rs.client.Subscribe(context.Background(), "CustomerRegistrationStarted").Channel()
}
