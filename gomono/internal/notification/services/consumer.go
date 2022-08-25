package notification

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

type CustomerRegistrationStartedEvent struct {
	ID       int64  `json:"id"`
	Identity string `json:"identity"`
	Customer struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	} `json:"customer"`
	Authentications []struct {
		ID        int64     `json:"id"`
		Code      string    `json:"code"`
		CreatedAt time.Time `json:"created_at"`
		ExpiredAt time.Time `json:"expired_at"`
	}
}

func consumeCustomerRegistrationStarted(ch <-chan *redis.Message) {
	for message := range ch {
		registrationEvent := CustomerRegistrationStartedEvent{}
		if err := json.Unmarshal([]byte(message.Payload), &registrationEvent); err != nil {
			logrus.Errorf("%w", err)
			continue
		}

		logrus.Infof("%+v", registrationEvent)

		// TODO: handle registration event
	}
}
