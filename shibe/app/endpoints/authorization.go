package shibeendpoints

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
)

const defaultTokenType = "bareer"

type AuthorizationRequest struct {
	Token string
	GrantType string
}

type AuthorizationResponse struct {
	AccessToken string
	TokenType string
	ExpiresIn int64
}

func (ep *defaultEndpoint) auth(c echo.Context) error {
	var req AuthorizationRequest

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		return err
	}

	if req.GrantType != defaultTokenType {
		return errors.New("No type found ")
	}

	resp := AuthorizationResponse{
		AccessToken: ep.config.AccessToken(),
		TokenType: defaultTokenType,
		ExpiresIn: 86400,
	}

	return ep.responseSuccess(c, resp)
}