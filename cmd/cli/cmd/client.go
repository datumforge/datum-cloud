package datumcloud

import (
	"context"
	"net/url"

	"github.com/datumforge/datum-cloud/internal/client"
)

// SetupClient will setup the datum cloud client
func SetupClient(ctx context.Context, host string) (client.Client, error) {
	config := client.NewDefaultConfig()

	opt, err := configureClientEndpoints(host)
	if err != nil {
		return nil, err
	}

	return client.New(config, opt)
}

// configureClientEndpoints will setup the base URL for the datum client
func configureClientEndpoints(datumCloudHost string) (opt client.ClientOption, err error) {
	if datumCloudHost == "" {
		return
	}

	baseURL, err := url.Parse(datumCloudHost)
	if err != nil {
		return nil, err
	}

	return client.WithBaseURL(baseURL), nil
}
