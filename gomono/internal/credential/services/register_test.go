package credential

import (
	"context"
	"reflect"
	"testing"

	"github.com/alicebob/miniredis/v2"
	repositories "github.com/bastianrob/gomono/internal/credential/repositories"
	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/bastianrob/gomono/pkg/schema"
	"github.com/go-redis/redis/v9"
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

func (m *RegistrationMock) Phone() string {
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
	CredentialRepository
}

func (m *CredentialRepositoryMock) FindCredentialByIdentity(ctx context.Context, identity string) (*repositories.FindCredentialByIdentityResult, error) {
	args := m.Called(ctx, identity)
	return args.Get(0).(*repositories.FindCredentialByIdentityResult), args.Error(1)
}

func (m *CredentialRepositoryMock) CountCredentialByIdentity(ctx context.Context, identity string) (int64, error) {
	args := m.Called(ctx, identity)
	return int64(args.Int(0)), args.Error(1)
}

func (m *CredentialRepositoryMock) CreateNewCustomer(ctx context.Context, input schema.CustomerRegisterInput) (*repositories.CreateNewCustomerMutationResult, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*repositories.CreateNewCustomerMutationResult), args.Error(1)
}

func newRegistrationMock(name, identity, phone, password, confirmation, provider string) *RegistrationMock {
	m := &RegistrationMock{}

	m.On("Name").Return(name)
	m.On("Identity").Return(identity)
	m.On("Phone").Return(phone)
	m.On("Password").Return(password)
	m.On("Confirmation").Return(confirmation)
	m.On("Provider").Return(provider)
	return m
}

func TestCredentialService_NewCustomer(t *testing.T) {
	credentialService := NewCredentialService(
		func() *CredentialRepositoryMock {
			m := &CredentialRepositoryMock{}
			m.On("CountCredentialByIdentity", mock.Anything, "exists@email.com").Return(1, nil)
			m.On("CountCredentialByIdentity", mock.Anything, mock.Anything).Return(0, nil)
			m.On("CreateNewCustomer", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(&repositories.CreateNewCustomerMutationResult{Credential: schema.Credential{ID: 1}}, nil)
			return m
		}(),
		nil,
	)

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
				reg: newRegistrationMock("someone", "someone@email.com", "+6282312345678", "passwd12345678", "passwd12345678", "email"),
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
				reg: newRegistrationMock("someone", "", "+6282312345678", "passwd12345678", "passwd12345678", "email"),
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
				reg: newRegistrationMock("someone", "i.am.not.an.email.com", "+6282312345678", "passwd12345678", "passwd12345678", "email"),
			},
			should: resp{
				returnError:  true,
				returnedID:   -1,
				errorMessage: "Please input a valid email address",
			},
		},
		{
			it:    "Should not be able to register when phone is empty",
			given: credentialService,
			when: args{
				ctx: context.Background(),
				reg: newRegistrationMock("someone", "someone@email.com", "", "passwd12345678", "passwd12345678", "email"),
			},
			should: resp{
				returnError:  true,
				returnedID:   -1,
				errorMessage: "Please input a valid phone number",
			},
		},
		{
			it:    "Should not be able to register when phone is not in e164 format",
			given: credentialService,
			when: args{
				ctx: context.Background(),
				reg: newRegistrationMock("someone", "someone@email.com", "453546", "passwd12345678", "passwd12345678", "email"),
			},
			should: resp{
				returnError:  true,
				returnedID:   -1,
				errorMessage: "Please input a valid phone number",
			},
		},
		{
			it:    "Should not be able to register when name is less than 2 characters",
			given: credentialService,
			when: args{
				ctx: context.Background(),
				reg: newRegistrationMock("n", "someone@email.com", "+6282312345678", "passwd12345678", "passwd12345678", "email"),
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
				reg: newRegistrationMock("someone", "someone@email.com", "+6282312345678", "passw<8", "passw<8", "email"),
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
				reg: newRegistrationMock("someone", "someone@email.com", "+6282312345678", "passwd12345678", "passwdDoesntmatch", "email"),
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
				reg: newRegistrationMock("someone", "exists@email.com", "+6282312345678", "passwd12345678", "passwd12345678", "email"),
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
				assert.Error(t, err, "should return an error")
				if err != nil {
					assert.Equal(t, tt.should.errorMessage, err.(*exception.Exception).Message, "error message does not match")
				}
			} else {
				assert.Nil(t, err, "should not return an error")
				assert.Equal(
					t, tt.should.returnedID,
					reflect.ValueOf(got).FieldByName("ID").Int(),
					"returned id does not match",
				)
			}
		})
	}
}

func TestCredentialService_publishCustomerRegistrationStartedEvent(t *testing.T) {
	miniredisServer := miniredis.RunT(t)
	type args struct {
		ctx    context.Context
		result *repositories.CreateNewCustomerMutationResult
	}
	type should struct {
		returnError bool
	}
	tests := []struct {
		it     string
		given  *CredentialService
		when   args
		should should
	}{{
		it: "Should successfully publish customer registration started event",
		given: NewCredentialService(nil, redis.NewClient(&redis.Options{
			Addr: miniredisServer.Addr(),
		})),
		when: args{
			ctx: context.Background(),
			result: &repositories.CreateNewCustomerMutationResult{
				Credential: schema.Credential{
					ID: 1000,
				},
			},
		},
		should: should{
			returnError: false,
		},
	}}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			err := tt.given.publishCustomerRegistrationStartedEvent(tt.when.ctx, tt.when.result)
			if tt.should.returnError {
				assert.Error(t, err, "should return an error")
			}

			assert.NoError(t, err, "should not return error")
		})

	}
}
