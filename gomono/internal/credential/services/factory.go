package credential

import (
	"context"

	repositories "github.com/bastianrob/gomono/internal/credential/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type CredentialRepository interface {
	FindCredentialByIdentity(context.Context, string) (*repositories.FindCredentialByIdentityResult, error)
	CountCredentialByIdentity(context.Context, string) (int64, error)
	CreateNewCustomer(ctx context.Context, identity, password, provider string) (*repositories.CreateNewCustomerMutationResult, error)
}

type CredentialService struct {
	repo     CredentialRepository
	validate *validator.Validate
}

func InitializeService() *CredentialService {
	repo := repositories.InitializeRepository()
	return NewCredentialService(repo)
}

func NewCredentialService(repo CredentialRepository) *CredentialService {
	return &CredentialService{
		repo:     repo,
		validate: validator.New(),
	}
}

type HasuraCustomClaims struct {
	AllowedRoles []string `json:"x-hasura-allowed-roles"`
	DefaultRole  string   `json:"x-hasura-default-role"`
	UserID       string   `json:"x-hasura-user-id"`
	PartnerID    int64    `json:"x-hasura-partner-id,omitempty"`
}

type StandardClaims struct {
	jwt.RegisteredClaims
	Claims HasuraCustomClaims `json:"https://hasura.io/jwt/claims"`
}
