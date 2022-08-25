package credential

import (
	"context"
	"time"

	"github.com/bastianrob/gomono/internal/credential/configs"
	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/machinebox/graphql"
)

type Credential struct {
	ID              int64            `json:"id"`
	Identity        string           `json:"identity"`
	Customer        Customer         `json:"customer"`
	Authentications []Authentication `json:"authentications"`
}

type Authentication struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CreateNewCustomerMutationResult struct {
	Credential Credential `json:"credential"`
}

func (repo *CredentialRepository) CreateNewCustomer(
	ctx context.Context,
	name,
	identity,
	phone,
	password,
	provider,
	authCode string,
	authExpiry time.Time,
) (*CreateNewCustomerMutationResult, error) {
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
	query.Var("name", name)
	query.Var("identity", identity)
	query.Var("phone", phone)
	query.Var("password", password)
	query.Var("provider", provider)
	query.Var("authCode", authCode)
	query.Var("authExpiry", authExpiry)

	resp := &CreateNewCustomerMutationResult{}
	if err := repo.gqlClient.Run(ctx, query, resp); err != nil {
		return nil, err
	}

	if resp.Credential.ID <= 0 {
		return nil, exception.New(nil, "Credential Not Found", exception.CodeNotFound)
	}

	return resp, nil
}
