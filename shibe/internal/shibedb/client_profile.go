package shibedb

import shibemodels "github.com/bosszanahub/shibe/internal/shibedb/models"

func (c *defaultClient) GetProfile(email string) shibemodels.Profile {
	return shibemodels.Profile{
		Email: "test",
		Name: "test",
		Age: 22,
	}
}