package credential

import (
	"net/http"

	"github.com/bastianrob/gomono/pkg/global"
	"github.com/labstack/echo/v4"
)

type HealthcheckResponse struct {
}

func (cont *CredentialController) Healthcheck(e echo.Context) error {
	response := global.ResponseDTO[HealthcheckResponse]{
		OK:      true,
		Message: "HC",
		Data:    HealthcheckResponse{},
	}

	e.JSON(http.StatusOK, response)
	return nil
}
