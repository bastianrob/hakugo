package credential

import (
	"context"
	"time"

	"github.com/bastianrob/gomono/pkg/global/idgen"
	"github.com/bastianrob/gomono/pkg/schema"
)

// Resend verification email
func (svc *CredentialService) Resend(ctx context.Context, email string) (any, error) {
	res, err := svc.repo.FindCredentialByIdentity(ctx, email)
	if err != nil {
		return nil, err
	}

	cred := res.Credential[0]
	input := &schema.InsertAuthenticationInput{
		CredentialID: cred.ID,
		Code:         idgen.AlphaNum20U(),
		ExpiredAt:    time.Now().Add(24 * time.Hour),
	}

	auth, err := svc.repo.CreateNewAuthentication(ctx, input)

	svc.publishVerificationEmailCommand(ctx, map[string]any{
		"name":          email,
		"email":         email,
		"code":          auth.Code,
		"redirect_host": "http://localhost",
	})

	return auth, err
}
