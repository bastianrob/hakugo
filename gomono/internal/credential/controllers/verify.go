package credential

import (
	"encoding/json"
	"net/http"

	"github.com/bastianrob/gomono/pkg/global"
	"github.com/labstack/echo/v4"
)

type AuthenticationVerifyRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (cont *CredentialController) AuthenticationVerify(e echo.Context) error {
	defer e.Request().Body.Close()

	payload := &global.RequestDTO[AuthenticationVerifyRequest]{}
	err := json.NewDecoder(e.Request().Body).Decode(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide a valid auth verification format")
	}

	token, err := cont.service.Verify(e.Request().Context(), payload.Data.Email, payload.Data.Code)
	if err != nil || token == "" {
		return echo.ErrUnauthorized
	}

	return e.JSON(http.StatusOK, AuthResponse{
		AccessToken: token,
	})
}
