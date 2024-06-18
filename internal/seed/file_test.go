package seed

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadCSVFile(t *testing.T) {
	// Create a temporary CSV file for testing
	tempFile, err := os.CreateTemp("", "test.csv")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Write some test data to the temporary CSV file
	data := []string{"1,John,Doe", "2,Jane,Smith"}

	err = os.WriteFile(tempFile.Name(), []byte(strings.Join(data, "\n")), 0600)
	require.NoError(t, err)

	// Call the loadCSVFile function
	upload, err := loadCSVFile(tempFile.Name())
	require.NoError(t, err)

	// Assert the returned values
	assert.NotNil(t, upload)
	assert.NotNil(t, upload.File)

	assert.Equal(t, tempFile.Name(), upload.Filename)
	assert.Equal(t, "text/csv", upload.ContentType)
}

func TestGetColumnIndex(t *testing.T) {
	headers := []string{"ID", "Name", "Age"}

	// Test when the column exists
	columnIndex := getColumnIndex(headers, "Name")
	assert.Equal(t, 1, columnIndex)

	// Test when the column does not exist
	columnIndex = getColumnIndex(headers, "Email")
	assert.Equal(t, -1, columnIndex)
}

func TestReadCSVFile(t *testing.T) {
	// Create a temporary CSV file for testing
	tempFile, err := os.CreateTemp("", "test.csv")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Write some test data to the temporary CSV file
	data := []string{"ID,FirstName,LastName", "1,John,Doe", "2,Jane,Smith"}

	err = os.WriteFile(tempFile.Name(), []byte(strings.Join(data, "\n")), 0600)
	require.NoError(t, err)

	// Call the readCSVFile function
	rows, err := readCSVFile(tempFile.Name())
	require.NoError(t, err)

	// Assert the returned values
	expectedRows := [][]string{{"ID", "FirstName", "LastName"}, {"1", "John", "Doe"}, {"2", "Jane", "Smith"}}
	assert.Equal(t, expectedRows, rows)
}
