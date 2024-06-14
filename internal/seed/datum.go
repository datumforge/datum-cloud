package seed

import (
	"context"
	"encoding/csv"
	"fmt"
	"time"

	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/models"
)

// Config represents provides the datum client and configuration for the seed client
type Client struct {
	*datumclient.DatumClient
	config *Config
}

// NewDefaultClient creates a new datum client using the default configuration variables
func NewDefaultClient() (*Client, error) {
	config, err := NewDefaultConfig()
	if err != nil {
		return nil, err
	}

	datumClient, err := config.createClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		DatumClient: datumClient,
		config:      config,
	}, nil
}

// NewClient creates a new datum client using the provided configuration variables
func (c *Config) NewClient() (*Client, error) {
	datumClient, err := c.createClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		DatumClient: datumClient,
		config:      c,
	}, nil
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

// AuthorizeOrganizationOnPAT authorizes the organization id on the personal access token id
func (c *Client) AuthorizeOrganizationOnPAT(ctx context.Context, orgID, patID string) error {
	input := datumclient.UpdatePersonalAccessTokenInput{
		AddOrganizationIDs: []string{orgID},
	}

	if _, err := c.UpdatePersonalAccessToken(ctx, patID, input); err != nil {
		return err
	}

	return nil
}

// GenerateSeedAPIToken generates an API token for the organization id to use for seeding
// and authenticates the client with the token for future requests
func (c *Client) GenerateSeedAPIToken(ctx context.Context, orgID string) error {
	expiresAt := time.Now().Add(time.Hour)

	input := datumclient.CreateAPITokenInput{
		Name:      fmt.Sprintf("seed token %s", orgID),
		OwnerID:   &orgID,
		ExpiresAt: &expiresAt, // expires in 1 hour
		Scopes:    []string{"read", "write"},
	}

	token, err := c.CreateAPIToken(ctx, input)
	if err != nil {
		return err
	}

	// Use the token to authenticate
	c.config.Token = token.CreateAPIToken.APIToken.Token

	// create a new client with the new token
	c.DatumClient, err = c.config.createClient()
	if err != nil {
		return err
	}

	return nil
}

// LoadOrganizations loads the organizations from the organizations.csv file
func (c *Client) LoadOrganizations(ctx context.Context) (string, error) {
	file := c.config.getOrgFilePath()

	upload, err := loadCSVFile(file)
	if err != nil {
		return "", err
	}

	org, err := c.CreateBulkCSVOrganization(ctx, upload)
	if err != nil {
		return "", err
	}

	// get the first org, this is the root org
	if len(org.CreateBulkCSVOrganization.Organizations) > 0 {
		return org.CreateBulkCSVOrganization.Organizations[0].ID, nil
	}

	return "", nil
}

// LoadGroups loads the groups from the groups.csv file
func (c *Client) LoadGroups(ctx context.Context) error {
	file := c.config.getGroupFilePath()

	upload, err := loadCSVFile(file)
	if err != nil {
		return err
	}

	if _, err := c.CreateBulkCSVGroup(ctx, upload); err != nil {
		return err
	}

	return nil
}

// LoadInvites loads the invites from the invites.csv file
func (c *Client) LoadInvites(ctx context.Context) error {
	file := c.config.getInviteFilePath()

	upload, err := loadCSVFile(file)
	if err != nil {
		return err
	}

	if _, err := c.CreateBulkCSVInvite(ctx, upload); err != nil {
		return err
	}

	return nil
}

// LoadSubscribers loads the subscribers from the subscribers.csv file
func (c *Client) LoadSubscribers(ctx context.Context) error {
	file := c.config.getSubscriberFilePath()

	upload, err := loadCSVFile(file)
	if err != nil {
		return err
	}

	if _, err := c.CreateBulkCSVSubscriber(ctx, upload); err != nil {
		return err
	}

	return nil
}

// RegisterUsers registers the users from the users.csv file
func (c *Client) RegisterUsers(ctx context.Context) error {
	file := c.config.getUserFilePath()

	upload, err := loadCSVFile(file)
	if err != nil {
		return err
	}

	reader := csv.NewReader(upload.File)
	records, _ := reader.ReadAll()

	for i, record := range records {
		// skip header row
		if i == 0 {
			continue
		}

		req := models.RegisterRequest{
			Email:     record[2],
			Password:  record[3],
			FirstName: record[0],
			LastName:  record[1],
		}

		if _, err := c.Register(ctx, &req); err != nil {
			return err
		}
	}

	return nil
}
