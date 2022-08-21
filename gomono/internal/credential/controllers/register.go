package credential

import (
	"encoding/json"
	"net/http"

	"github.com/bastianrob/gomono/pkg/global"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Fullname string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Pass     string `json:"password"`
	Conf     string `json:"confirmation"`
	Prov     string `json:"provider"`
}

func (r *RegisterRequest) Name() string {
	return r.Fullname
}

func (r *RegisterRequest) Identity() string {
	return r.Email
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

type RegisterResponse struct {
	ID int64 `json:"id"`
}

func (cont *CredentialController) Register(e echo.Context) error {
	defer e.Request().Body.Close()

	payload := &global.RequestDTO[RegisterRequest]{}
	err := json.NewDecoder(e.Request().Body).Decode(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide a valid registration format")
	}

	id, err := cont.service.NewCustomer(e.Request().Context(), &payload.Data)
	if err != nil || id <= 0 {
		return echo.ErrInternalServerError
	}

	return e.JSON(http.StatusCreated, RegisterResponse{
		ID: id,
	})
}
