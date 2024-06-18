package seed

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserEmails(t *testing.T) {
	filename := "users.csv"

	t.Run("File does not exist, generate all random emails", func(t *testing.T) {
		numUsers := 10

		emails, err := getUserEmails(filename, numUsers)
		require.NoError(t, err)

		assert.Len(t, emails, numUsers)
	})

	t.Run("File exists, num users is less than users in file", func(t *testing.T) {
		numUsers := 2

		// Create a temporary file with test data
		file, err := os.CreateTemp("", "users.csv")
		require.NoError(t, err)

		defer os.Remove(file.Name())

		// Write test data to the file
		data := []string{"ID,FirstName,LastName,Email", "1,John,Doe,john@example.com", "2,Jane,Smith,jane@example.com", "3,Meow,Cat,meow@example.com"}

		err = os.WriteFile(file.Name(), []byte(strings.Join(data, "\n")), 0600)
		require.NoError(t, err)

		emails, err := getUserEmails(file.Name(), numUsers)
		require.NoError(t, err)

		assert.Len(t, emails, numUsers)

		// Check that the emails are the same as the first n emails in the test data
		for i := range emails {
			assert.Contains(t, data[i+1], emails[i])
		}
	})

	t.Run("File exists, num users is greater than users in file", func(t *testing.T) {
		numUsers := 4

		// Create a temporary file with test data
		file, err := os.CreateTemp("", "users.csv")
		require.NoError(t, err)

		defer os.Remove(file.Name())

		// Write test data to the file
		data := []string{"ID,FirstName,LastName,Email", "1,John,Doe,john@example.com", "2,Jane,Smith,jane@example.com", "3,Meow,Cat,meow@example.com"}

		err = os.WriteFile(file.Name(), []byte(strings.Join(data, "\n")), 0600)
		require.NoError(t, err)

		emails, err := getUserEmails(file.Name(), numUsers)
		require.NoError(t, err)

		assert.Len(t, emails, numUsers)

		// Check that the emails are the same as the first n emails in the test data
		for i := range emails {
			if i <= len(data)-2 { // skip the header row
				assert.Contains(t, data[i+1], emails[i])
			}
		}
	})
}
