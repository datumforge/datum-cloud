package datumcloud

import (
	"github.com/datumforge/datum-cloud/internal/seed"
)

// NewSeedClient creates a new datum Seed client, requiring a token to be set
func NewSeedClient() (*seed.Client, error) {
	conf, err := seed.NewDefaultConfig()
	if err != nil {
		return nil, err
	}

	if Config.String("directory") != "" {
		conf.Directory = Config.String("directory")
	}

	if Config.String("token") == "" {
		return nil, ErrDatumAPITokenMissing
	}

	conf.Token = Config.String("token")

	return conf.NewClient()
}
