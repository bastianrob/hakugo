package credential

import (
	"context"

	"github.com/bastianrob/gomono/internal/credential/configs"
	repositories "github.com/bastianrob/gomono/internal/credential/repositories"
	"github.com/bastianrob/gomono/pkg/schema"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v9"
	"github.com/golang-jwt/jwt/v4"
)

type CredentialRepository interface {
	FindCredentialByIdentity(context.Context, string) (*repositories.FindCredentialByIdentityResult, error)
	CountCredentialByIdentity(context.Context, string) (int64, error)
	CreateNewCustomer(ctx context.Context, input schema.CustomerRegisterInput) (*repositories.CreateNewCustomerMutationResult, error)
	FindAuthenticationByCode(ctx context.Context, code string) (*schema.Authentication, error)
	SetAuthenticationAsUsed(ctx context.Context, authID, credentialID int64) (*schema.Authentication, error)
	CreateNewAuthentication(ctx context.Context, input *schema.InsertAuthenticationInput) (*schema.Authentication, error)
}

type CredentialService struct {
	repo        CredentialRepository
	redisClient redis.Cmdable
	validate    *validator.Validate
}

func InitializeService() *CredentialService {
	return NewCredentialService(
		repositories.InitializeRepository(),
		redis.NewClient(&redis.Options{
			Addr:     configs.App.Redis.Host,
			Password: configs.App.Redis.Pass,
			DB:       configs.App.Redis.DB,
		}),
	)
}

func NewCredentialService(repo CredentialRepository, redisClient *redis.Client) *CredentialService {
	return &CredentialService{
		repo:        repo,
		redisClient: redisClient,
		validate:    validator.New(),
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
