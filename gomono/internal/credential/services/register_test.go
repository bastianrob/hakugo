package credential

import (
	"context"
	"testing"

	repositories "github.com/bastianrob/gomono/internal/credential/repositories"
	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RegistrationMock struct {
	mock.Mock
	Registration
}

func (m *RegistrationMock) Name() string {
	args := m.Called()
	return args.String(0)
}

func (m *RegistrationMock) Identity() string {
	args := m.Called()
	return args.String(0)
}

func (m *RegistrationMock) Password() string {
	args := m.Called()
	return args.String(0)
}

func (m *RegistrationMock) Confirmation() string {
	args := m.Called()
	return args.String(0)
}

func (m *RegistrationMock) Provider() string {
	args := m.Called()
	return args.String(0)
}

type CredentialRepositoryMock struct {
	mock.Mock
}

func (m *CredentialRepositoryMock) FindCredentialByIdentity(ctx context.Context, identity string) (*repositories.FindCredentialByIdentityResult, error) {
	args := m.Called(ctx, identity)
	return args.Get(0).(*repositories.FindCredentialByIdentityResult), args.Error(1)
}

func (m *CredentialRepositoryMock) CountCredentialByIdentity(ctx context.Context, identity string) (int64, error) {
	args := m.Called(ctx, identity)
	return int64(args.Int(0)), args.Error(1)
}

func (m *CredentialRepositoryMock) CreateNewCustomer(ctx context.Context, identity string, password string, provider string) (*repositories.CreateNewCustomerMutationResult, error) {
	args := m.Called(ctx, identity, password, provider)
	return args.Get(0).(*repositories.CreateNewCustomerMutationResult), args.Error(1)
}

func newRegistrationMock(name, identity, password, confirmation, provider string) *RegistrationMock {
	m := &RegistrationMock{}

	m.On("Name").Return(name)
	m.On("Identity").Return(identity)
	m.On("Password").Return(password)
	m.On("Confirmation").Return(confirmation)
	m.On("Provider").Return(provider)
	return m
}

func TestCredentialService_NewCustomer(t *testing.T) {
	credentialService := NewCredentialService(func() *CredentialRepositoryMock {
		m := &CredentialRepositoryMock{}
		m.On("CountCredentialByIdentity", mock.Anything, "exists@email.com").Return(1, nil)
		m.On("CountCredentialByIdentity", mock.Anything, mock.Anything).Return(0, nil)
		m.On("CreateNewCustomer", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&repositories.CreateNewCustomerMutationResult{Insertion: struct {
			ID int64 "json:\"id\""
		}{ID: 1}}, nil)
		return m
	}())

	type args struct {
		ctx context.Context
		reg Registration
	}
	type resp struct {
		returnedID   int64
		returnError  bool
		errorMessage string
	}
	tests := []struct {
		it     string
		given  *CredentialService
		when   args
		should resp
	}{
		{
			it:    "Should succesfully registered",
			given: credentialService,
			when: args{
				ctx: context.Background(),
				reg: newRegistrationMock("someone", "someone@email.com", "passwd12345678", "passwd12345678", "email"),
			},
			should: resp{
				returnError: false,
				returnedID:  1,
			},
		},
		{
			it:    "Should not be able to register when email is empty",
			given: credentialService,
			when: args{
				ctx: context.Background(),
				reg: newRegistrationMock("someone", "", "passwd12345678", "passwd12345678", "email"),
			},
			should: resp{
				returnError:  true,
				returnedID:   -1,
				errorMessage: "Please input a valid email address",
			},
		},
		{
			it:    "Should not be able to register when email is not valid",
			given: credentialService,
			when: args{
				ctx: context.Background(),
				reg: newRegistrationMock("someone", "i.am.not.an.email.com", "passwd12345678", "passwd12345678", "email"),
			},
			should: resp{
				returnError:  true,
				returnedID:   -1,
				errorMessage: "Please input a valid email address",
			},
		},
		{
			it:    "Should not be able to register when name is less than 2 characters",
			given: credentialService,
			when: args{
				ctx: context.Background(),
				reg: newRegistrationMock("n", "someone@email.com", "passwd12345678", "passwd12345678", "email"),
			},
			should: resp{
				returnError:  true,
				returnedID:   -1,
				errorMessage: "Name should be longer than 2 characters",
			},
		},
		{
			it:    "Should not be able to register when password is less than 8 char",
			given: credentialService,
			when: args{
				ctx: context.Background(),
				reg: newRegistrationMock("someone", "someone@email.com", "passw<8", "passw<8", "email"),
			},
			should: resp{
				returnError:  true,
				returnedID:   -1,
				errorMessage: "Password should be at least 8 characters long",
			},
		},
		{
			it:    "Should not be able to register when confirmation password does not match",
			given: credentialService,
			when: args{
				ctx: context.Background(),
				reg: newRegistrationMock("someone", "someone@email.com", "passwd12345678", "passwdDoesntmatch", "email"),
			},
			should: resp{
				returnError:  true,
				returnedID:   -1,
				errorMessage: "Confirmation password doesn't match",
			},
		},
		{
			it:    "Should not be able to register when email already exists",
			given: credentialService,
			when: args{
				ctx: context.Background(),
				reg: newRegistrationMock("someone", "exists@email.com", "passwd12345678", "passwd12345678", "email"),
			},
			should: resp{
				returnError:  true,
				returnedID:   -1,
				errorMessage: "Email address: exists@email.com is already registered",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			got, err := tt.given.NewCustomer(tt.when.ctx, tt.when.reg)
			if tt.should.returnError {
				assert.Error(t, err, "Should return an error")
				if err != nil {
					assert.Equal(t, tt.should.errorMessage, err.(*exception.Exception).Message, "error message does not match")
				}
			}

			assert.Equal(t, tt.should.returnedID, got, "returned id does not match")
		})
	}
}
