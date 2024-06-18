package client

import (
	"context"
	"net/http"

	"github.com/datumforge/datum/pkg/httpsling"

	"github.com/datumforge/datum-cloud/internal/v1/models"
)

// Client is the interface that wraps the DatumCloud API REST client methods
type Client interface {
	// WorkspaceCreate creates an organizational hierarchy for a workspace
	WorkspaceCreate(context.Context, *models.WorkspaceRequest) (*models.WorkspaceReply, error)
}

// NewWithDefaults creates a new API v1 client with default configuration
func NewWithDefaults() (Client, error) {
	conf := NewDefaultConfig()

	return New(conf)
}

// New creates a new API v1 client that implements the Client interface
func New(config Config, opts ...ClientOption) (_ Client, err error) {
	c := &APIv1{
		Config: config,
	}

	// apply the options to the client
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	// create the HTTP sling client if it is not set
	if c.HTTPSlingClient == nil {
		c.HTTPSlingClient, err = newHTTPClient(c.Config)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// newHTTPClient creates a new HTTP sling client with the given configuration
func newHTTPClient(config Config) (*httpsling.Client, error) {
	// copy the values from the base config to the httpsling config
	if config.HTTPSling == nil {
		config.HTTPSling = &httpsling.Config{}
	}

	if config.HTTPSling.BaseURL == "" {
		config.HTTPSling.BaseURL = config.BaseURL.String()
	}

	client := httpsling.Create(config.HTTPSling)

	return client, nil
}

// APIv1 implements the Client interface and provides methods to interact with the Datum Cloud API
type APIv1 struct {
	// Config is the configuration for the APIv1 client
	Config Config
	// HTTPSlingClient is the HTTP client for the APIv1 client
	HTTPSlingClient *httpsling.Client
}

// Ensure the APIv1 implements the Client interface
var _ Client = &APIv1{}

// WorkspaceCreate creates an organizational hierarchy for a new workspace based on the name, environment(s), bucket(s), and
// relationship(s) provided in the request
func (c *APIv1) WorkspaceCreate(ctx context.Context, in *models.WorkspaceRequest) (out *models.WorkspaceReply, err error) {
	req := c.HTTPSlingClient.NewRequestBuilder(http.MethodPost, "/v1/workspace")
	req.Body(in)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	if err := resp.ScanJSON(&out); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, newRequestError(resp.StatusCode(), out.Error)
	}

	return out, nil
}
