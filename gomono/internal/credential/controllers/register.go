package credential

import (
	"encoding/json"
	"net/http"

	credential "github.com/bastianrob/gomono/internal/credential/services"
	"github.com/bastianrob/gomono/pkg/global"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	User string `json:"identity"`
	Pass string `json:"password"`
	Conf string `json:"confirmation"`
	Prov string `json:"provider"`
}

func (r *RegisterRequest) Identity() string {
	return r.User
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

func (r *RegisterRequest) SetPassword(pwd string) credential.Registration {
	r.Pass = pwd
	return r
}

func (r *RegisterRequest) SetProvider(prov string) credential.Registration {
	r.Prov = prov
	return r
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

	response := &global.ResponseDTO[RegisterResponse]{
		OK:      false,
		Message: "Authenticated",
		Data: RegisterResponse{
			ID: id,
		},
	}

	return e.JSON(http.StatusCreated, response)
}
