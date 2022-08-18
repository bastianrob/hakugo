package credential

import (
	"context"
	"testing"

	"github.com/bastianrob/gomono/internal/credential/configs"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
)

func TestCredentialService_findCredentialByIdentity(t *testing.T) {
	configs.App.GraphQL.AuthHeader = "x-hasura-admin-secret"
	configs.App.GraphQL.AuthSecret = "12345678"

	client := graphql.NewClient("http://localhost/v1/graphql")
	service := NewCredentialService(client)
	service.findCredentialByIdentity(context.Background(), "someone@email.com")
}

func TestCredentialService_countCredentialByIdentity(t *testing.T) {
	configs.App.GraphQL.AuthHeader = "x-hasura-admin-secret"
	configs.App.GraphQL.AuthSecret = "12345678"

	client := graphql.NewClient("http://localhost/v1/graphql")
	service := NewCredentialService(client)
	result, err := service.countCredentialByIdentity(context.Background(), "someone@email.com")
	assert.NoError(t, err)
	assert.Equal(t, true, result > 0, result)
}
