package seed

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	orgsFileName = "orgs.csv"
)

// getOrgFilePath returns the full path to the orgs file
func (c *Config) getOrgFilePath() string {
	return fmt.Sprintf("%s/%s", c.Directory, orgsFileName)
}

// generateOrgData generates org data and writes it to a CSV file
func (c *Config) generateOrgData() error {
	if c.NumOrganizations <= 0 {
		return nil
	}

	file, err := os.Create(c.getOrgFilePath())
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	// Add column headers
	if err := csvWriter.Write([]string{"Name"}); err != nil {
		return err
	}

	// Add data
	for i := 0; i < c.NumOrganizations; i++ {
		if err := csvWriter.Write([]string{
			cases.Title(language.English, cases.Compact).String(gofakeit.Company()),
		}); err != nil {
			return err
		}
	}

	// Flush the data to the file
	csvWriter.Flush()

	return nil
}
