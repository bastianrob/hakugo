package credential

import (
	"context"

	"github.com/bastianrob/gomono/internal/credential/configs"
	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/bastianrob/gomono/pkg/schema"
	"github.com/machinebox/graphql"
)

type CreateNewCustomerMutationResult struct {
	Credential schema.Credential `json:"credential"`
}

func (repo *CredentialRepository) CreateNewCustomer(ctx context.Context, input schema.CustomerRegisterInput) (*CreateNewCustomerMutationResult, error) {
	query := graphql.NewRequest(`
	mutation createNewCustomer(
		$name: String!, $identity: String!, $phone: String!, $password: String!, $provider: String!,
		$authCode: bpchar!, $authExpiry: timestamptz!
	) {
		credential: insert_credential_one(
			object: {
				identity: $identity,
				password: $password,
				provider: $provider,
				banned: false
				customer: {
					data: {
						name: $name,
						email: $identity,
						phone: $phone
					}
				}
				authentications: {
					data: [{
						code: $authCode,
						expired_at: $authExpiry,
						used: false
					}]
				}
			}
		) {
			id
			identity
			customer {
				name
				email
				phone
			}
      authentications(limit: 1) {
        code
        expired_at
      }
		}
	}`)

	query.Header.Add(configs.App.GraphQL.AuthHeader, configs.App.GraphQL.AuthSecret)
	query.Var("name", input.Name)
	query.Var("identity", input.Identity)
	query.Var("phone", input.Phone)
	query.Var("password", input.Password)
	query.Var("provider", input.Provider)
	query.Var("authCode", input.AuthCode)
	query.Var("authExpiry", input.AuthExpiry)

	resp := &CreateNewCustomerMutationResult{}
	if err := repo.gqlClient.Run(ctx, query, resp); err != nil {
		return nil, err
	}

	if resp.Credential.ID <= 0 {
		return nil, exception.New(nil, "Credential Not Found", exception.CodeNotFound)
	}

	return resp, nil
}

func (repo *CredentialRepository) SetAuthenticationAsUsed(ctx context.Context, authID int64) (*schema.Authentication, error) {
	query := graphql.NewRequest(`
	mutation updateAuthentication($id: bigint!) {
		authentication: update_authentication_by_pk(
			pk_columns: { id: $id },
			_set: { used: true }
		) {
			used
		}
	}`)

	query.Header.Add(configs.App.GraphQL.AuthHeader, configs.App.GraphQL.AuthSecret)
	query.Var("id", authID)

	resp := &struct {
		Authentication *schema.Authentication `json:"authentication"`
	}{}
	if err := repo.gqlClient.Run(ctx, query, resp); err != nil || resp.Authentication == nil || !resp.Authentication.Used {
		return nil, exception.New(err, "Failed to updated authentication code", exception.CodeUnexpectedError)
	}

	return resp.Authentication, nil
}
