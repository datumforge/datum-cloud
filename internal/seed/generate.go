package seed

import "os"

func (c *Config) GenerateData() error {
	// Create the directory if it doesn't exist
	mode := os.ModeDir | 0755 //nolint:mnd
	if _, err := os.Stat(c.Directory); os.IsNotExist(err) {
		if err := os.Mkdir(c.Directory, mode); err != nil {
			return err
		}
	}

	// Generate the organization data
	if err := c.generateOrgData(); err != nil {
		return err
	}

	// Generate the group data
	if err := c.generateGroupData(); err != nil {
		return err
	}

	// Generate the user data
	if err := c.generateUserData(); err != nil {
		return err
	}

	// Generate the invite data
	if err := c.generateInviteData(); err != nil {
		return err
	}

	// Generate the subscriber data
	return c.generateSubscriberData()
}

func (c *Config) GenerateUserData() error {
	// Generate the user data
	return c.generateUserData()
}
