package app

import (
	shibeendpoints "github.com/bosszanahub/shibe/app/endpoints"
	shibeentities "github.com/bosszanahub/shibe/app/entities"
	shibeusecases "github.com/bosszanahub/shibe/app/usecases"
	"github.com/labstack/echo/v4"
)

type App interface {}

func New(e *echo.Echo, cfg shibeentities.Config) App {

	tokenValidate := shibeusecases.NewTokenValidateUseCase(cfg.AccessToken())

	ep := shibeendpoints.NewEndPoint(cfg, tokenValidate)
	ep.EnableEndpoint(e)

	return ep
}