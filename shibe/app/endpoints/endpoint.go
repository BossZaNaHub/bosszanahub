package shibeendpoints

import (
	shibeentities "github.com/bosszanahub/shibe/app/entities"
	shibeusecases "github.com/bosszanahub/shibe/app/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Endpoint interface {
	EnableEndpoint(ec *echo.Echo)
}

type defaultEndpoint struct {
	config shibeentities.Config
	tkValidate shibeusecases.TokenValidateUseCase
}

func NewEndPoint(config shibeentities.Config, tkValidate shibeusecases.TokenValidateUseCase) Endpoint {
	return &defaultEndpoint{
		config: config,
		tkValidate: tkValidate,
	}
}

func (ep *defaultEndpoint) EnableEndpoint(ec *echo.Echo) {
	/* Authenticate */
}

func (ep *defaultEndpoint) createErrorResponse() shibeentities.ErrorResponse {
	resp := shibeentities.ErrorResponse{
		Success: false,
		ErrorCode: 500,
		Message: "",
	}

	return resp
}

func (ep *defaultEndpoint) responseError(c echo.Context) error {
	resp := ep.createErrorResponse()

	return c.JSON(http.StatusBadRequest, resp)
}

func (ep *defaultEndpoint) responseSuccess(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, result)
}

func (ep *defaultEndpoint) GetheaderToken(c echo.Context) string {
	bearerToken := c.Request().Header.Get("Authorization")
	token := strings.Split(bearerToken, " ")
	if len(token) != 2 {
		return ""
	}
	return token[1]
}

func (ep *defaultEndpoint) ValidateToken(c echo.Context) error {
	token := ep.GetheaderToken(c)

	err := ep.tkValidate.ValidateAccessToken(token)

	return err
}