package seed

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v7"
)

const (
	usersFileName = "users.csv"
)

// getUserFilePath returns the full path to the users file
func (c *Config) getUserFilePath() string {
	return fmt.Sprintf("%s/%s", c.Directory, usersFileName)
}

// generateUserData generates user data and writes it to a CSV file
func (c *Config) generateUserData() error {
	if c.NumUsers <= 0 {
		return nil
	}

	file, err := os.Create(c.getUserFilePath())
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	// Add column headers
	headers := []string{"First Name", "Last Name", "Email", "Password", "AuthProvider", "OrganizationIDs"}
	if err := csvWriter.Write(headers); err != nil {
		return err
	}

	// Add data
	for i := 0; i < c.NumUsers; i++ {
		if err := csvWriter.Write(
			generateUserDetails(),
		); err != nil {
			return err
		}
	}

	// Flush the data to the file
	csvWriter.Flush()

	return nil
}

// generateUserDetails generates user details using the gofakeit library
func generateUserDetails() []string {
	p := gofakeit.Person()
	passwordLength := 20

	return []string{
		p.FirstName,
		p.LastName,
		p.Contact.Email,
		// there is not guarantee that the password will have special characters, so we add one to ensure
		fmt.Sprintf("%s!", gofakeit.Password(true, true, true, true, false, passwordLength)),
		"CREDENTIALS",
		"[ORGANIZATION_ID]",
	}
}
