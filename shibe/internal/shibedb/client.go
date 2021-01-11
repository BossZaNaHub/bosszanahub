package shibedb

import shibemodels "github.com/bosszanahub/shibe/internal/shibedb/models"

type Client interface {
	GetProfile(email string) shibemodels.Profile
}