package credential

import (
	"context"
	"fmt"
	"time"
)

func (svc *CredentialService) Authenticate(ctx context.Context, identity, password, role string) (string, error) {
	result, err := svc.repo.FindCredentialByIdentity(ctx, identity)
	if err != nil {
		return "", err
	}

	credential := result.Credential[0]
	if !isPasswordMatch(credential.Password, password) {
		return "", err
	}

	token, err := generateAuthToken(
		time.Now(),
		fmt.Sprintf("%d", credential.ID),
		crawlPartnerID(result),
		identity,
		role,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
