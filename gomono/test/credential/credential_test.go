package credential_test

import (
	"context"
	"os"
	"testing"

	"github.com/bastianrob/gomono/internal/credential/configs"
	repositories "github.com/bastianrob/gomono/internal/credential/repositories"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	configs.Init()
	code := m.Run()
	os.Exit(code)
}

func TestCredentialService_findCredentialByIdentity(t *testing.T) {
	configs.App.GraphQL.AuthHeader = "x-hasura-admin-secret"
	configs.App.GraphQL.AuthSecret = "12345678"

	client := graphql.NewClient("http://localhost:8080/v1/graphql")
	repository := repositories.NewCredentialRepository(client)
	// service := credential.NewCredentialService(repository, nil)
	res, err := repository.FindCredentialByIdentity(context.Background(), "mail@mail.mail")
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Credential)
}

func TestCredentialService_countCredentialByIdentity(t *testing.T) {
	configs.App.GraphQL.AuthHeader = "x-hasura-admin-secret"
	configs.App.GraphQL.AuthSecret = "12345678"

	client := graphql.NewClient("http://localhost:8080/v1/graphql")
	repository := repositories.NewCredentialRepository(client)
	// service := credential.NewCredentialService(repository, nil)
	res, err := repository.CountCredentialByIdentity(context.Background(), "someone@email.com")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), res)
}
