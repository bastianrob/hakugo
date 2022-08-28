package credential

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/bastianrob/gomono/pkg/global"
	"github.com/bastianrob/gomono/pkg/global/idgen"
	"github.com/bastianrob/gomono/pkg/schema"
	"github.com/sirupsen/logrus"
)

type Registration interface {
	Name() string
	Identity() string
	Phone() string
	Password() string
	Confirmation() string
	Provider() string
}

// NewCustomer creates a new customer and returns its ID, otherwise returns -1 and error
func (svc *CredentialService) NewCustomer(ctx context.Context, reg Registration) (any, error) {
	err := svc.validateRegistration(ctx, reg)
	if err != nil {
		return nil, err
	}

	authCode := idgen.AlphaNum20U()
	authExpiry := time.Now().Add(24 * time.Hour)
	hashedPassowrd, _ := hashPassword(reg.Password())
	result, err := svc.repo.CreateNewCustomer(ctx, schema.CustomerRegisterInput{
		Name:       reg.Name(),
		Identity:   reg.Identity(),
		Phone:      reg.Phone(),
		Provider:   reg.Provider(),
		Password:   hashedPassowrd,
		AuthCode:   authCode,
		AuthExpiry: authExpiry,
	})
	if err != nil {
		logrus.Error(err)
		return nil, exception.New(err, "Failed to create new customer", exception.CodeUnexpectedError)
	}

	svc.publishVerificationEmailCommand(ctx, map[string]any{
		"name":          result.Credential.Customer.Name,
		"email":         result.Credential.Customer.Email,
		"code":          result.Credential.Authentications[0].Code,
		"redirect_host": "http://localhost",
	})

	return result.Credential, nil
}

func (svc *CredentialService) validateRegistration(ctx context.Context, reg Registration) error {
	isPasswordMatchConfirmation := reg.Password() == reg.Confirmation()
	if !isPasswordMatchConfirmation {
		return exception.New(nil, "Confirmation password doesn't match", exception.CodeValidationError)
	}

	count, err := svc.repo.CountCredentialByIdentity(ctx, reg.Identity())
	isAlreadyExists := count >= 1 || err != nil
	if isAlreadyExists {
		return exception.New(err, fmt.Sprintf("Email address: %s is already registered", reg.Identity()), exception.CodeValidationError)
	}

	if err = svc.validate.Var(reg.Name(), "required,min=2,max=128"); err != nil {
		return exception.New(err, "Name should be longer than 2 characters", exception.CodeValidationError)
	}

	if err = svc.validate.Var((reg.Identity()), "required,email"); err != nil {
		return exception.New(err, "Please input a valid email address", exception.CodeValidationError)
	}

	if err = svc.validate.Var(reg.Phone(), "required,e164"); err != nil {
		return exception.New(err, "Please input a valid phone number", exception.CodeValidationError)
	}

	if err = svc.validate.Var(reg.Password(), "required,min=8"); err != nil {
		return exception.New(err, "Password should be at least 8 characters long", exception.CodeValidationError)
	}

	if err = svc.validate.Var(reg.Provider(), "required,oneof=email google github facebook"); err != nil {
		return exception.New(err, "Provider must be one of email, google, github, or facebook", exception.CodeValidationError)
	}

	return nil
}

func (svc *CredentialService) publishVerificationEmailCommand(
	ctx context.Context,
	data map[string]any,
) error {

	if svc.redisClient == nil || reflect.ValueOf(svc.redisClient).IsNil() {
		return nil
	}

	attempt := 0
	var err error
	for {
		attempt += 1
		err = svc.redisClient.
			// TODO: Redis will broadcast to all subs which is dangerous if subscriber have multiple pods
			Publish(ctx, "SendVerificationEmailCommand", global.EventDTO[any]{
				IssuedAt: time.Now(),
				Type:     "SendVerificationEmailCommand",
				Data:     data,
			}).
			Err()

		if err == nil || attempt >= 5 {
			break
		}
	}

	return err
}
