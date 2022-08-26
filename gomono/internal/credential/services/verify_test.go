package credential

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialService_Verify(t *testing.T) {
	type args struct {
		ctx  context.Context
		code string
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
			given: NewCredentialService(nil, nil),
			when: args{
				ctx:  context.Background(),
				code: "CORRECT-CODE-4567890",
			},
			should: should{
				returnErr: false,
			},
		},
		{
			it:    "Should return error when code is NOT valid",
			given: NewCredentialService(nil, nil),
			when: args{
				ctx:  context.Background(),
				code: "INVALID-CODE-4567890",
			},
			should: should{
				returnErr: true,
			},
		},
		{
			it:    "Should return error when code is expired",
			given: NewCredentialService(nil, nil),
			when: args{
				ctx:  context.Background(),
				code: "EXPIRED-CODE-4567890",
			},
			should: should{
				returnErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			err := tt.given.Verify(tt.when.ctx, tt.when.code)
			if tt.should.returnErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
