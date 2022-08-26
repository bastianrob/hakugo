package notification

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/mailjet/mailjet-apiv3-go"
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
	} `json:"authentications"`
}

func (svc *NotificationService) consumeCustomerRegistrationStarted(ch <-chan *redis.Message) {
	for message := range ch {
		registrationEvent := CustomerRegistrationStartedEvent{}
		if err := json.Unmarshal([]byte(message.Payload), &registrationEvent); err != nil {
			logrus.Errorf("%s", err)
			continue
		}

		buffer := bytes.Buffer{}
		mailVerificationTemplate.Execute(&buffer, MailVerificationTemplate{
			Name:  registrationEvent.Customer.Name,
			Email: registrationEvent.Customer.Email,
			Code:  registrationEvent.Authentications[0].Code,
			Host:  "http://localhost",
		})

		mailBody := strings.TrimSpace(strings.ReplaceAll(buffer.String(), "\t", ""))
		messagesInfo := []mailjet.InfoMessagesV31{
			{
				From: &mailjet.RecipientV31{
					Email: "hello@bastianrob.xyz",
					Name:  "Robin's Email Bot",
				},
				To: &mailjet.RecipientsV31{
					mailjet.RecipientV31{
						Email: registrationEvent.Customer.Email,
						Name:  registrationEvent.Customer.Name,
					},
				},
				Subject:  "Verify Your Email",
				TextPart: mailBody,
			},
		}
		messages := &mailjet.MessagesV31{Info: messagesInfo}
		_, err := svc.mailjetClient.SendMailV31(messages)
		if err != nil {
			logrus.Error(err)
		}
	}
}
