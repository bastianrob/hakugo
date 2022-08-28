package notification

import (
	"context"

	"github.com/go-redis/redis/v9"
)

func (rs *RedisSubsriber) SendVerificationEmailCommand() <-chan *redis.Message {
	return rs.client.Subscribe(context.Background(), "SendVerificationEmailCommand").Channel()
}
