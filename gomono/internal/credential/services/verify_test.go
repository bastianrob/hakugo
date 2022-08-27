package credential

import (
	"context"
	"testing"
	"time"

	"github.com/bastianrob/gomono/internal/credential/configs"
	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/bastianrob/gomono/pkg/global"
	"github.com/bastianrob/gomono/pkg/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (m *CredentialRepositoryMock) FindAuthenticationByCode(ctx context.Context, code string) (*schema.Authentication, error) {
	args := m.Called(ctx, code)
	auth, ok := args.Get(0).(*schema.Authentication)
	if ok {
		return auth, args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *CredentialRepositoryMock) SetAuthenticationAsUsed(ctx context.Context, authID int64) (*schema.Authentication, error) {
	args := m.Called(ctx, authID)
	auth, ok := args.Get(0).(*schema.Authentication)
	if ok {
		return auth, args.Error(1)
	}

	return nil, args.Error(1)
}

func TestCredentialService_Verify(t *testing.T) {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	yesterday := now.AddDate(0, 0, -1)
	configs.Init()
	global.Init()

	credentialService := NewCredentialService(
		func() *CredentialRepositoryMock {
			m := &CredentialRepositoryMock{}
			m.On("FindAuthenticationByCode", mock.Anything, "CORRECT-CODE-4567890").Return(&schema.Authentication{ExpiredAt: tomorrow}, nil)
			m.On("FindAuthenticationByCode", mock.Anything, "EXPIRED-CODE-4567890").Return(&schema.Authentication{ExpiredAt: yesterday}, nil)
			m.On("FindAuthenticationByCode", mock.Anything, "USEDALR-CODE-4567890").Return(&schema.Authentication{ExpiredAt: tomorrow, Used: true}, nil)
			m.On("FindAuthenticationByCode", mock.Anything, "INVALID-CODE-4567890").Return(nil, exception.New(nil, "NOT FOUND", exception.CodeNotFound))
			m.On("SetAuthenticationAsUsed", mock.Anything, mock.Anything).Return(&schema.Authentication{}, nil)
			return m
		}(),
		nil,
	)

	type args struct {
		ctx   context.Context
		code  string
		email string
	}
	type should struct {
		returnErr bool
	}
	tests := []struct {
		it     string
		given  *CredentialService
		when   args
		should should
	}{
		{
			it:    "Should not return error when code is valid",
			given: credentialService,
			when: args{
				ctx:   context.Background(),
				code:  "CORRECT-CODE-4567890",
				email: "mail@mail.mail",
			},
			should: should{
				returnErr: false,
			},
		},
		{
			it:    "Should return error when code is NOT valid",
			given: credentialService,
			when: args{
				ctx:   context.Background(),
				code:  "INVALID-CODE-4567890",
				email: "mail@mail.mail",
			},
			should: should{
				returnErr: true,
			},
		},
		{
			it:    "Should return error when code is expired",
			given: credentialService,
			when: args{
				ctx:   context.Background(),
				code:  "EXPIRED-CODE-4567890",
				email: "mail@mail.mail",
			},
			should: should{
				returnErr: true,
			},
		},
		{
			it:    "Should return error when code is used already",
			given: credentialService,
			when: args{
				ctx:   context.Background(),
				code:  "USEDALR-CODE-4567890",
				email: "mail@mail.mail",
			},
			should: should{
				returnErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			token, err := tt.given.Verify(tt.when.ctx, tt.when.email, tt.when.code)
			if tt.should.returnErr {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}
