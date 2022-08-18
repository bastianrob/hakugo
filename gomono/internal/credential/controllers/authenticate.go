package credential

import (
	"encoding/json"
	"net/http"

	"github.com/bastianrob/gomono/pkg/global"
	"github.com/labstack/echo/v4"
)

type AuthRequest struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
}

func (cont *CredentialController) Authenticate(e echo.Context) error {
	defer e.Request().Body.Close()

	payload := &global.RequestDTO[AuthRequest]{}
	err := json.NewDecoder(e.Request().Body).Decode(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide a valid credential format")
	}

	token, err := cont.service.Authenticate(
		e.Request().Context(),
		payload.Data.Identity,
		payload.Data.Password,
		payload.Data.Role,
	)

	if err != nil || token == "" {
		return echo.ErrUnauthorized
	}

	return e.JSON(http.StatusOK, AuthResponse{
		AccessToken: token,
	})
}
