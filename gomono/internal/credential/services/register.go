package credential

import (
	"context"
	"fmt"
	"reflect"
	"time"

	repositories "github.com/bastianrob/gomono/internal/credential/repositories"
	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/bastianrob/gomono/pkg/global"
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

	hashedPassowrd, _ := hashPassword(reg.Password())
	result, err := svc.repo.CreateNewCustomer(ctx, reg.Name(), reg.Identity(), reg.Phone(), hashedPassowrd, reg.Provider())
	if err != nil {
		logrus.Error(err)
		return nil, exception.New(err, "Failed to create new customer", exception.CodeUnexpectedError)
	}

	svc.publishCustomerRegistrationStartedEvent(ctx, result)
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

func (svc *CredentialService) publishCustomerRegistrationStartedEvent(
	ctx context.Context,
	result *repositories.CreateNewCustomerMutationResult,
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
			Publish(ctx, "CustomerRegistrationStarted", global.EventDTO[any]{
				IssuedAt: time.Now(),
				Type:     "CustomerRegistrationStarted",
				Data:     result.Credential,
			}).
			Err()

		if err == nil || attempt >= 5 {
			break
		}
	}

	return err
}
