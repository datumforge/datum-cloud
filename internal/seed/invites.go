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
	emails, err := getUserEmails(c.getUserFilePath(), c.NumInvites)
	if err != nil {
		return err
	}

	for _, email := range emails {
		if err := csvWriter.Write([]string{email, getRole()}); err != nil {
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

// getUserEmail returns a subset of user emails from the users.csv file
// if the file does not exist, it will return a random emails instead
func getUserEmails(filename string, numUsers int) ([]string, error) {
	emails := []string{}

	// if the file does not exist, generate random emails
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		for range numUsers {
			emails = append(emails, gofakeit.Email())
		}

		return emails, nil
	}

	// get the emails from the file
	records, err := readCSVFile(filename)
	if err != nil {
		return nil, err
	}

	// find the email column
	emailIndex := getColumnIndex(records[0], "Email")
	if emailIndex == -1 {
		return nil, fmt.Errorf("%w: %s", ErrColumnNotFound, "Email")
	}

	// make sure we don't go out of bounds
	userCount := len(records) - 1
	generateAdditionalUsers := 0

	if numUsers > userCount {
		generateAdditionalUsers = numUsers - userCount
	}

	userEmails := min(numUsers, userCount)

	// grab the first userCount emails
	records = records[1 : userEmails+1]
	for _, record := range records {
		emails = append(emails, record[emailIndex])
	}

	// generate additional users if needed
	for range generateAdditionalUsers {
		emails = append(emails, gofakeit.Email())
	}

	return emails, nil
}
