package seed

import (
	"github.com/mcuadros/go-defaults"
)

// Config is the configuration for the seed package
type Config struct {
	// Directory is the directory to save generated data
	Directory string `json:"directory" koanf:"directory" default:"demodata"`
	// Token is the token to use for the datum client
	Token string `json:"token" koanf:"token" default:""`
	// NumOrganizations is the number of organizations to generate
	NumOrganizations int `json:"numOrganizations" koanf:"numOrganizations" default:"1"`
	// NumUsers is the number of users to generate
	NumUsers int `json:"NumUsers" koanf:"NumUsers" default:"10"`
	// NumGroups is the number of groups to generate
	NumGroups int `json:"NumGroups" koanf:"NumGroups" default:"10"`
	// NumInvites is the number of invites to generate
	NumInvites int `json:"NumInvites" koanf:"NumInvites" default:"5"`
}

// NewDefaultConfig returns a new Config with default values
func NewDefaultConfig() (*Config, error) {
	// Set default values
	conf := &Config{}
	defaults.SetDefaults(conf)

	return conf, nil
}
