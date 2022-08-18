package credential

import (
	"context"

	"github.com/bastianrob/gomono/internal/credential/configs"
	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/machinebox/graphql"
)

type CreateNewCustomerMutationResult struct {
	Insertion struct {
		ID int64 `json:"id"`
	} `json:"insertion"`
}

func (svc *CredentialService) createNewCustomer(ctx context.Context, reg Registration) (*CreateNewCustomerMutationResult, error) {
	query := graphql.NewRequest(`
		mutation createNewCustomer($identity: String!, $password: String!, $provider: String!) {
			insertion: insert_credential_one(object: {identity: $identity, password: $password, provider: $provider}) {
				id
			}
		}
	`)

	query.Header.Add(configs.App.GraphQL.AuthHeader, configs.App.GraphQL.AuthSecret)
	query.Var("identity", reg.Identity())
	query.Var("password", reg.Password())
	query.Var("provider", reg.Provider())

	resp := &CreateNewCustomerMutationResult{}
	if err := svc.gqlClient.Run(ctx, query, resp); err != nil {
		return nil, err
	}

	if resp.Insertion.ID <= 0 {
		return nil, exception.New(nil, "Credential Not Found")
	}

	return resp, nil
}
