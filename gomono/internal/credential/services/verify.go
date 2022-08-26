package credential

import (
	"context"
	"time"

	"github.com/bastianrob/gomono/pkg/exception"
)

// Verify authentication code
func (svc *CredentialService) Verify(ctx context.Context, code string) error {
	// 1. Find the code
	auth, err := svc.repo.FindAuthenticationByCode(ctx, code)
	if exc, isException := exception.IsException(err); isException {
		return exc
	} else if err != nil {
		return exception.New(err, "Unknown error occurred", exception.CodeUnexpectedError)
	}

	// 2. Validate whether it's used or expired
	now := time.Now()
	if auth.Used {
		return exception.New(nil, "Code is already used", exception.CodeValidationError)
	} else if now.Unix() > auth.ExpiredAt.Unix() {
		return exception.New(nil, "Code is already expired", exception.CodeValidationError)
	}

	// 3. Set the auth as used
	// TODO svc.repo.SetAuthenticationAsUsed(ctx, auth.ID)

	return nil
}
