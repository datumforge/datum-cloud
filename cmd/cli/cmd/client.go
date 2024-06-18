package datumcloud

import (
	"context"
	"net/url"

	"github.com/datumforge/datum-cloud/internal/client"
)

// SetupClient will setup the datum cloud client
func SetupClient(ctx context.Context) (client.Client, error) {
	config := client.NewDefaultConfig()

	opt, err := configureClientEndpoints()
	if err != nil {
		return nil, err
	}

	return client.New(config, opt)
}

// configureClientEndpoints will setup the base URL for the datum client
func configureClientEndpoints() (client.ClientOption, error) {
	baseURL, err := url.Parse(DatumCloudHost)
	if err != nil {
		return nil, err
	}

	return client.WithBaseURL(baseURL), nil
}
