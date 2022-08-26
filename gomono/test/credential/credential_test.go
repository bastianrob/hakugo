package credential_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/bastianrob/gomono/internal/credential/configs"
	repositories "github.com/bastianrob/gomono/internal/credential/repositories"
	"github.com/bastianrob/gomono/pkg/schema"
	"github.com/google/go-cmp/cmp"
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

func TestCredentialService_FindAuthenticationByCode(t *testing.T) {
	configs.App.GraphQL.AuthHeader = "x-hasura-admin-secret"
	configs.App.GraphQL.AuthSecret = "12345678"

	client := graphql.NewClient("http://localhost:8080/v1/graphql")
	repository := repositories.NewCredentialRepository(client)
	// service := credential.NewCredentialService(repository, nil)
	res, err := repository.FindAuthenticationByCode(context.Background(), "12345123456789067890")
	assert.NoError(t, err)
	assert.True(t, cmp.Equal(res, &schema.Authentication{
		ID:        1,
		Code:      "12345123456789067890",
		CreatedAt: time.Date(2022, time.August, 25, 15, 3, 56, 588894000, time.UTC),
		ExpiredAt: time.Date(2022, time.August, 26, 15, 3, 16, 0, time.UTC),
		Used:      false,
	}))
}
