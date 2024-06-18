package datum

import "github.com/datumforge/datum/pkg/datumclient"

// NewDefaultClient creates a new datum client using the default configuration variables
func NewDefaultClient() (*datumclient.DatumClient, error) {
	config, err := NewDefaultConfig()
	if err != nil {
		return nil, err
	}

	return config.createClient()
}

// NewClient creates a new datum client using the provided configuration variables
func (c *Config) NewClient() (*datumclient.DatumClient, error) {
	return c.createClient()
}

// CreateDatumClient creates a new datum client using the DATUM_TOKEN configuration variable
func (c *Config) createClient() (*datumclient.DatumClient, error) {
	if c.Token == "" {
		return nil, ErrAPITokenMissing
	}

	config := datumclient.NewDefaultConfig()

	opt := datumclient.WithCredentials(datumclient.Authorization{
		BearerToken: c.Token})

	return datumclient.New(config, opt)
}
