package credential

import (
	"github.com/bastianrob/gomono/internal/credential/configs"
	"github.com/machinebox/graphql"
)

type CredentialRepository struct {
	gqlClient *graphql.Client
}

func InitializeRepository() *CredentialRepository {
	return NewCredentialRepository(graphql.NewClient(configs.App.GraphQL.Host))
}

func NewCredentialRepository(gqlClient *graphql.Client) *CredentialRepository {
	return &CredentialRepository{
		gqlClient: gqlClient,
	}
}
