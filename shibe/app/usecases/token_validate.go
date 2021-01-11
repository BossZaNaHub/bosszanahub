package shibeusecases

import "errors"

type TokenValidateUseCase interface {
	ValidateAccessToken(token string) error
}

type defaultTokenValidateUseCase struct {
	accessToken string
}

func NewTokenValidateUseCase(token string) TokenValidateUseCase {
	return &defaultTokenValidateUseCase{
		accessToken: token,
	}
}

func (u *defaultTokenValidateUseCase) ValidateAccessToken(token string) error {
	if u.accessToken != token {
		return nil
	}

	return errors.New("Invalid Token ")
}