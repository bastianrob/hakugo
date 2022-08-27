package credential

import (
	"context"

	"github.com/bastianrob/gomono/internal/credential/configs"
	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/bastianrob/gomono/pkg/schema"
	"github.com/machinebox/graphql"
)

type FindCredentialByIdentityResult struct {
	Credential []schema.Credential `json:"credential"`
}

// findCredentialByIdentity return error if not found
func (repo *CredentialRepository) FindCredentialByIdentity(ctx context.Context, identity string) (*FindCredentialByIdentityResult, error) {
	query := graphql.NewRequest(`
		query findCredentialByIdentity($identity: String!) {
			credential(where: {identity: {_eq: $identity}}) {
				id
				identity
				password
				banned
				partners(limit: 1) {
					id
				}
			}
		}
	`)

	query.Header.Add(configs.App.GraphQL.AuthHeader, configs.App.GraphQL.AuthSecret)
	query.Var("identity", identity)

	resp := &FindCredentialByIdentityResult{}
	if err := repo.gqlClient.Run(ctx, query, resp); err != nil {
		return nil, err
	}

	if len(resp.Credential) <= 0 {
		return nil, exception.New(nil, "Credential Not Found", exception.CodeNotFound)
	}

	return resp, nil
}

type CountCredentialByIdentityResult struct {
	Credential struct {
		Aggregate struct {
			Count int64 `json:"count"`
		} `json:"aggregate"`
	} `json:"credential"`
}

// countCredentialByIdentity return error 0 if not found
func (repo *CredentialRepository) CountCredentialByIdentity(ctx context.Context, identity string) (int64, error) {
	query := graphql.NewRequest(`
		query countCredentialByIdentity($identity: String!) {
			credential: credential_aggregate(where: {identity: {_eq: $identity}}) {
				aggregate {
					count
				}
			}
		}
	`)

	query.Header.Add(configs.App.GraphQL.AuthHeader, configs.App.GraphQL.AuthSecret)
	query.Var("identity", identity)

	resp := &CountCredentialByIdentityResult{}
	if err := repo.gqlClient.Run(ctx, query, resp); err != nil {
		return 0, err
	}

	return resp.Credential.Aggregate.Count, nil
}

// FindAuthenticationByCode returns error if not found
func (repo *CredentialRepository) FindAuthenticationByCode(ctx context.Context, code string) (*schema.Authentication, error) {
	query := graphql.NewRequest(`
		query findAuthenticationByCode($where: authentication_bool_exp) {
      	authentications: authentication(limit: 1, where: $where) {
				id
				credential_id
				created_at
				expired_at
				code
				used
			}
		}	
	`)

	query.Header.Add(configs.App.GraphQL.AuthHeader, configs.App.GraphQL.AuthSecret)
	query.Var("where", map[string]any{
		"code": map[string]any{
			"_eq": code,
		},
	})

	resp := &struct {
		Authentications []schema.Authentication `json:"authentications"`
	}{}
	if err := repo.gqlClient.Run(ctx, query, resp); err != nil {
		return nil, err
	}

	if len(resp.Authentications) <= 0 {
		return nil, exception.New(nil, "Authentication code does not exists", exception.CodeNotFound)
	}

	return &resp.Authentications[0], nil
}
