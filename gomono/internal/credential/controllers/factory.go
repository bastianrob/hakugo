package credential

import (
	"context"

	credential "github.com/bastianrob/gomono/internal/credential/services"
	"github.com/labstack/echo/v4"
)

type CredentialService interface {
	Authenticate(ctx context.Context, identity, password, role string) (string, error)
	NewCustomer(ctx context.Context, reg credential.Registration) (int64, error)
}

type CredentialController struct {
	service CredentialService
}

func InitializeController() *CredentialController {
	service := credential.InitializeService()
	return NewCredentialController(service)
}

func NewCredentialController(service CredentialService) *CredentialController {
	return &CredentialController{
		service: service,
	}
}

func (cont *CredentialController) Routes(e *echo.Echo) {
	e.GET("/", cont.Healthcheck)
	e.POST("/authenticate", cont.Authenticate)
	e.POST("/register", cont.Register)
}
