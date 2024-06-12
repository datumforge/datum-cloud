package seed

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"

	"github.com/brianvoe/gofakeit/v7"
)

const (
	invitesFileName = "invites.csv"
)

// getInviteFilePath returns the full path to the invites file
func (c *Config) getInviteFilePath() string {
	return fmt.Sprintf("%s/%s", c.Directory, invitesFileName)
}

// generateInviteData generates invite data and writes it to a CSV file
func (c *Config) generateInviteData() error {
	if c.NumInvites <= 0 {
		return nil
	}

	file, err := os.Create(c.getInviteFilePath())
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	// Add column headers
	if err := csvWriter.Write([]string{"Recipient", "Role"}); err != nil {
		return err
	}

	// Add data
	for i := 0; i < c.NumInvites; i++ {
		if err := csvWriter.Write([]string{gofakeit.Email(), getRole()}); err != nil {
			return err
		}
	}

	// Flush the data to the file
	csvWriter.Flush()

	return nil
}

var (
	validRoles = []string{"MEMBER", "ADMIN"}
)

// getRole returns a random role from the validRoles slice
func getRole() string {
	return validRoles[rand.Intn(len(validRoles))] //nolint:gosec
}
