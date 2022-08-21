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

func (repo *CredentialRepository) CreateNewCustomer(
	ctx context.Context,
	name,
	identity,
	phone,
	password,
	provider string,
) (*CreateNewCustomerMutationResult, error) {
	query := graphql.NewRequest(`
		mutation createNewCustomer($name: String!, $identity: String!, $phone: String!, $password: String!, $provider: String!) {
			insertion: insert_credential_one(object: {identity: $identity, password: $password, provider: $provider, customers: {data: [{name: $name, email: $identity, phone: $phone}]}}) {
				id
			}
		}
	`)

	query.Header.Add(configs.App.GraphQL.AuthHeader, configs.App.GraphQL.AuthSecret)
	query.Var("name", name)
	query.Var("identity", identity)
	query.Var("phone", phone)
	query.Var("password", password)
	query.Var("provider", provider)

	resp := &CreateNewCustomerMutationResult{}
	if err := repo.gqlClient.Run(ctx, query, resp); err != nil {
		return nil, err
	}

	if resp.Insertion.ID <= 0 {
		return nil, exception.New(nil, "Credential Not Found")
	}

	return resp, nil
}
