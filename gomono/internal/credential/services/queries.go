package credential

import (
	"context"

	"github.com/bastianrob/gomono/internal/credential/configs"
	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/machinebox/graphql"
)

type FindCredentialByIdentityResult struct {
	Credential []struct {
		ID       int64  `json:"id"`
		Identity string `json:"identity"`
		Password string `json:"password"`
		Banned   bool   `json:"banned"`
		Partners []struct {
			ID int64 `json:"id"`
		} `json:"partners"`
	} `json:"credential"`
}

// findCredentialByIdentity return error if not found
func (svc *CredentialService) findCredentialByIdentity(ctx context.Context, identity string) (*FindCredentialByIdentityResult, error) {
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
	if err := svc.gqlClient.Run(ctx, query, resp); err != nil {
		return nil, err
	}

	if len(resp.Credential) <= 0 {
		return nil, exception.New(nil, "Credential Not Found")
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
func (svc *CredentialService) countCredentialByIdentity(ctx context.Context, identity string) (int64, error) {
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
	if err := svc.gqlClient.Run(ctx, query, resp); err != nil {
		return 0, err
	}

	return resp.Credential.Aggregate.Count, nil
}
