package client

import (
	"net/http"
	"net/url"

	"github.com/datumforge/datum/pkg/httpsling"
)

// Config is the configuration for the Datum Cloud API client
type Config struct {
	// BaseURL is the base URL for the Datum API
	BaseURL *url.URL `json:"baseUrl" yaml:"base_url" default:"http://localhost:17610"`
	// HTTPSling is the configuration for the HTTPSling client
	HTTPSling *httpsling.Config
}

// NewDefaultConfig returns a new default configuration for the Datum Cloud API client
func NewDefaultConfig() Config {
	return defaultClientConfig
}

var defaultClientConfig = Config{
	BaseURL: &url.URL{
		Scheme: "http",
		Host:   "localhost:17610",
	},
	HTTPSling: &httpsling.Config{
		Headers: &http.Header{
			"Accept":          []string{httpsling.ContentTypeJSONUTF8},
			"Accept-Language": []string{"en-US,en"},
			"Content-Type":    []string{httpsling.ContentTypeJSONUTF8},
		},
	},
}
