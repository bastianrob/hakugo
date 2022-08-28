package notification

import (
	"github.com/go-redis/redis/v9"
	"github.com/mailjet/mailjet-apiv3-go"
)

type Subscription interface {
	SendVerificationEmailCommand() <-chan *redis.Message
}

type NotificationService struct {
	subscription  Subscription
	mailjetClient *mailjet.Client
}

func NewNotificationService(
	subscription Subscription,
	mailjetClient *mailjet.Client,
) *NotificationService {
	svc := &NotificationService{
		subscription:  subscription,
		mailjetClient: mailjetClient,
	}

	return svc
}

func (svc *NotificationService) Run() {
	go svc.consumeVerificationEmailCommand(svc.subscription.SendVerificationEmailCommand())
}
