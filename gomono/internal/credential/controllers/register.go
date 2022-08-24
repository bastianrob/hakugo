package credential

import (
	"encoding/json"
	"net/http"

	"github.com/bastianrob/gomono/pkg/exception"
	"github.com/bastianrob/gomono/pkg/global"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Fullname string `json:"name"`
	Mail     string `json:"email"`
	Phon     string `json:"phone"`
	Pass     string `json:"password"`
	Conf     string `json:"confirmation"`
	Prov     string `json:"provider"`
}

func (r *RegisterRequest) Name() string {
	return r.Fullname
}

func (r *RegisterRequest) Identity() string {
	return r.Mail
}

func (r *RegisterRequest) Phone() string {
	return r.Phon
}

func (r *RegisterRequest) Password() string {
	return r.Pass
}

func (r *RegisterRequest) Confirmation() string {
	return r.Conf
}

func (r *RegisterRequest) Provider() string {
	return r.Prov
}

func (cont *CredentialController) Register(e echo.Context) error {
	defer e.Request().Body.Close()

	payload := &global.RequestDTO[RegisterRequest]{}
	err := json.NewDecoder(e.Request().Body).Decode(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide a valid registration format")
	}

	res, err := cont.service.NewCustomer(e.Request().Context(), &payload.Data)
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
