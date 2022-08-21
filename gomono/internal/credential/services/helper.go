package credential

import (
	"time"

	"github.com/bastianrob/gomono/internal/credential/configs"
	repositories "github.com/bastianrob/gomono/internal/credential/repositories"
	"github.com/bastianrob/gomono/pkg/global"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func generateAuthToken(now time.Time, userID string, partnerID int64, identity, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, StandardClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.Config.Name,
			Subject:   identity,
			Audience:  []string{},
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * 7 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        userID,
		},
		Claims: HasuraCustomClaims{
			AllowedRoles: []string{"customer", "partner", "anonymous"},
			DefaultRole:  role,
			UserID:       identity,
			PartnerID:    partnerID,
		},
	})

	signedToken, err := token.SignedString([]byte(configs.App.JWT.Secret))
	return signedToken, err
}

func hashPassword(input string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(hash), err
}

func isPasswordMatch(stored string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(stored), []byte(input))
	return err == nil
}

func crawlPartnerID(res *repositories.FindCredentialByIdentityResult) int64 {
	if len(res.Credential) <= 0 ||
		len(res.Credential[0].Partners) <= 0 ||
		res.Credential[0].Partners[0].ID <= 0 {
		return 0
	}

	return res.Credential[0].Partners[0].ID
}
