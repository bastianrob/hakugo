package credential

import (
	"context"
	"fmt"

	"github.com/bastianrob/gomono/pkg/exception"
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
func (svc *CredentialService) NewCustomer(ctx context.Context, reg Registration) (int64, error) {
	isPasswordMatchConfirmation := reg.Password() == reg.Confirmation()
	if !isPasswordMatchConfirmation {
		return -1, exception.New(nil, "Confirmation password doesn't match")
	}

	count, err := svc.repo.CountCredentialByIdentity(ctx, reg.Identity())
	isAlreadyExists := count >= 1 || err != nil
	if isAlreadyExists {
		return -1, exception.New(err, fmt.Sprintf("Email address: %s is already registered", reg.Identity()))
	}

	if err = svc.validate.Var(reg.Name(), "required,min=2,max=128"); err != nil {
		return -1, exception.New(err, "Name should be longer than 2 characters")
	}

	if err = svc.validate.Var((reg.Identity()), "required,email"); err != nil {
		return -1, exception.New(err, "Please input a valid email address")
	}

	if err = svc.validate.Var(reg.Phone(), "required,e164"); err != nil {
		return -1, exception.New(err, "Please input a valid phone number")
	}

	if err = svc.validate.Var(reg.Password(), "required,min=8"); err != nil {
		return -1, exception.New(err, "Password should be at least 8 characters long")
	}

	if err = svc.validate.Var(reg.Provider(), "required,oneof=email google github facebook"); err != nil {
		return -1, exception.New(err, "Provider must be one of email, google, github, or facebook")
	}

	hashedPassowrd, _ := hashPassword(reg.Password())
	result, err := svc.repo.CreateNewCustomer(ctx, reg.Name(), reg.Identity(), reg.Phone(), hashedPassowrd, reg.Provider())
	if err != nil {
		return -1, exception.New(err, "Failed to create new customer")
	}

	return result.Insertion.ID, nil
}
