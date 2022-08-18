package credential

import (
	"context"

	"github.com/bastianrob/gomono/pkg/exception"
)

type Registration interface {
	Identity() string
	Password() string
	Confirmation() string
	Provider() string

	SetPassword(string) Registration
	SetProvider(string) Registration
}

// NewCustomer creates a new customer and returns its ID, otherwise returns -1 and error
func (svc *CredentialService) NewCustomer(ctx context.Context, reg Registration) (int64, error) {
	isPasswordMatchConfirmation := reg.Password() == reg.Confirmation()
	if !isPasswordMatchConfirmation {
		return -1, exception.New(nil, "Confirmation password doesn't match")
	}

	count, err := svc.countCredentialByIdentity(ctx, reg.Identity())
	isAlreadyExists := count >= 1 || err != nil
	if isAlreadyExists {
		return -1, exception.New(nil, "Same identity already exists")
	}

	hashedPassowrd, _ := hashPassword(reg.Password())
	reg.SetPassword(hashedPassowrd)

	result, err := svc.createNewCustomer(ctx, reg)
	if err != nil {
		return -1, exception.New(err, "Failed to create new customer")
	}

	return result.Insertion.ID, nil
}
