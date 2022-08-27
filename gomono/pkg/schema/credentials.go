package schema

import "time"

type Credential struct {
	ID              int64            `json:"id,omitempty"`
	Identity        string           `json:"identity,omitempty"`
	Password        string           `json:"password,omitempty"`
	Banned          bool             `json:"banned,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Partners        []Partner        `json:"partners,omitempty"`
	Authentications []Authentication `json:"authentications,omitempty"`
}

type Authentication struct {
	ID           int64     `json:"id,omitempty"`
	CredentialID int64     `json:"credential_id,omitempty"`
	Code         string    `json:"code,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	ExpiredAt    time.Time `json:"expired_at,omitempty"`
	Used         bool      `json:"used,omitempty"`
}

type Customer struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type Partner struct {
	ID int64 `json:"id,omitempty"`
}

type CustomerRegisterInput struct {
	Name       string    `json:"name,omitempty"`
	Identity   string    `json:"identity,omitempty"`
	Phone      string    `json:"phone,omitempty"`
	Password   string    `json:"password,omitempty"`
	Provider   string    `json:"provider,omitempty"`
	AuthCode   string    `json:"auth_code,omitempty"`
	AuthExpiry time.Time `json:"auth_expiry,omitempty"`
}
