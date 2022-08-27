package credential

import (
	"context"
	"fmt"
	"time"

	"github.com/bastianrob/gomono/pkg/exception"
)

// Verify authentication code
// return authToken and error
func (svc *CredentialService) Verify(ctx context.Context, email, code string) (string, error) {
	// 1. Find the code
	auth, err := svc.repo.FindAuthenticationByCode(ctx, code)
	if exc, isException := exception.IsException(err); isException {
		return "", exc
	} else if err != nil {
		return "", exception.New(err, "Unknown error occurred", exception.CodeUnexpectedError)
	}

	// 2. Validate whether it's used or expired
	now := time.Now()
	if auth.Used {
		return "", exception.New(nil, "Code is already used", exception.CodeValidationError)
	} else if now.Unix() > auth.ExpiredAt.Unix() {
		return "", exception.New(nil, "Code is already expired", exception.CodeValidationError)
	}

	// 3. Set the auth as used
	if _, exc := svc.repo.SetAuthenticationAsUsed(ctx, auth.ID); err != nil {
		return "", exc
	}

	token, _ := generateAuthToken(
		time.Now(),
		fmt.Sprintf("%d", auth.CredentialID),
		0,
		email,
		"customer",
	)

	return token, nil
}
