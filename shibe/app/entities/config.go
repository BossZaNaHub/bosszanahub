package shibeentities

type Config struct {
	accessToken string
}

func NewConfig(accessToken string) Config {
	return Config{
		accessToken: accessToken,
	}
}

func (c Config) AccessToken() string {
	return c.accessToken
}
