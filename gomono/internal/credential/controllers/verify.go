package credential

import (
	"encoding/json"
	"net/http"

	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/bastianrob/gomono/pkg/global"
	"github.com/labstack/echo/v4"
)

type AuthenticationVerifyRequest struct {
	Email    string `json:"email,omitempty"`
	Code     string `json:"code,omitempty"`
	Activate bool   `json:"activate,omitempty"`
}

func (cont *CredentialController) AuthenticationVerify(e echo.Context) error {
	defer e.Request().Body.Close()

	payload := &global.RequestDTO[AuthenticationVerifyRequest]{}
	err := json.NewDecoder(e.Request().Body).Decode(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide a valid auth verification format")
	}

	token, err := cont.service.Verify(e.Request().Context(), payload.Data.Email, payload.Data.Code, payload.Data.Activate)
	if exc, isException := exception.IsException(err); isException {
		return e.JSON(http.StatusBadRequest, global.ErrorDTO{
			Message: exc.Message,
			Extensions: &global.ErrorExtension{
				Code:  exc.Code,
				Field: "$",
			},
		})
	} else if err != nil || token == "" {
		return echo.ErrUnauthorized
	}

	return e.JSON(http.StatusOK, AuthResponse{
		AccessToken: token,
	})
}
