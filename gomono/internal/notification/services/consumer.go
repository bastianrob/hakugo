package notification

import (
	"bytes"
	"encoding/json"

	"github.com/bastianrob/gomono/pkg/global"
	"github.com/go-redis/redis/v9"
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/sirupsen/logrus"
)

type SendVerificationEmailCommand struct {
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	Code         string `json:"code,omitempty"`
	RedirectHost string `json:"redirect_host,omitempty"`
}

func (svc *NotificationService) consumeVerificationEmailCommand(ch <-chan *redis.Message) {
	for message := range ch {
		registrationEvent := global.EventDTO[SendVerificationEmailCommand]{}
		if err := json.Unmarshal([]byte(message.Payload), &registrationEvent); err != nil {
			logrus.Errorf("%s", err)
			continue
		}

		buffer := bytes.Buffer{}
		mailVerificationTemplate.Execute(&buffer, MailVerificationTemplate{
			Name:  registrationEvent.Data.Name,
			Email: registrationEvent.Data.Email,
			Code:  registrationEvent.Data.Code,
			Host:  registrationEvent.Data.RedirectHost,
		})

		messagesInfo := []mailjet.InfoMessagesV31{
			{
				From: &mailjet.RecipientV31{
					Email: "hello@bastianrob.xyz",
					Name:  "Robin's Email Bot",
				},
				To: &mailjet.RecipientsV31{
					mailjet.RecipientV31{
						Email: "robin.bas90@gmail.com", //registrationEvent.Data.Email,
						Name:  registrationEvent.Data.Name,
					},
				},
				Subject:  "Verify Your Email",
				HTMLPart: buffer.String(),
			},
		}
		messages := &mailjet.MessagesV31{Info: messagesInfo}
		_, err := svc.mailjetClient.SendMailV31(messages)
		if err != nil {
			logrus.Error(err)
		}
	}
}
