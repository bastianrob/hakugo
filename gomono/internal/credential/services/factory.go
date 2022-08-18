package credential

import (
	"github.com/bastianrob/gomono/internal/credential/configs"
	"github.com/golang-jwt/jwt/v4"
	"github.com/machinebox/graphql"
)

type CredentialService struct {
	gqlClient *graphql.Client
}

func InitializeService() *CredentialService {
	return NewCredentialService(graphql.NewClient(configs.App.GraphQL.Host))
}

func NewCredentialService(gqlClient *graphql.Client) *CredentialService {
	return &CredentialService{
		gqlClient: gqlClient,
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
