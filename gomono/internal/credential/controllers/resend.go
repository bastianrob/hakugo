package credential

import (
	"encoding/json"
	"net/http"

	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/bastianrob/gomono/pkg/global"
	"github.com/labstack/echo/v4"
)

type ResendRequest struct {
	Email string `json:"email,omitempty"`
}

func (cont *CredentialController) Resend(e echo.Context) error {
	defer e.Request().Body.Close()

	payload := &global.RequestDTO[ResendRequest]{}
	err := json.NewDecoder(e.Request().Body).Decode(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide a valid registration format")
	}

	res, err := cont.service.Resend(e.Request().Context(), payload.Data.Email)
	if exc, ok := exception.IsException(err); ok {
		e.Logger().Error(exc.Message)
		return e.JSON(http.StatusBadRequest, global.ErrorDTO{
			Message: exc.Message,
			Extensions: &global.ErrorExtension{
				Code:  exc.Code,
				Field: "$",
			},
		})
	} else if err != nil {
		e.Logger().Error(err)
		return echo.ErrInternalServerError
	}

	return e.JSON(http.StatusCreated, res)
}
