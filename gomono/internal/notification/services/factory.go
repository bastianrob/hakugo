package notification

import (
	"github.com/go-redis/redis/v9"
	"github.com/mailjet/mailjet-apiv3-go"
)

type Subscription interface {
	CustomerRegistrationStarted() <-chan *redis.Message
}

type NotificationService struct {
	subscription  Subscription
	mailjetClient mailjet.ClientInterface
}

func NewNotificationService(
	subscription Subscription,
	mailjetClient mailjet.ClientInterface,
) *NotificationService {
	svc := &NotificationService{
		subscription:  subscription,
		mailjetClient: mailjetClient,
	}

	return svc
}

func (svc *NotificationService) Run() {
	go consumeCustomerRegistrationStarted(svc.subscription.CustomerRegistrationStarted())
}
