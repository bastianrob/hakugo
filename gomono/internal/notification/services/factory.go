package notification

import "github.com/go-redis/redis/v9"

type Subscription interface {
	CustomerRegistrationStarted() <-chan *redis.Message
}

type NotificationService struct {
	subscription Subscription
}

func NewNotificationService(subscription Subscription) *NotificationService {
	svc := &NotificationService{
		subscription: subscription,
	}

	return svc
}

func (svc *NotificationService) Run() {
	go consumeCustomerRegistrationStarted(svc.subscription.CustomerRegistrationStarted())
}
