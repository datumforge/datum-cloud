package seed

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	sliceutil "github.com/datumforge/datum/pkg/utils/slice"
)

const (
	groupsFileName = "groups.csv"
)

// getGroupFilePath returns the full path to the invites file
func (c *Config) getGroupFilePath() string {
	return fmt.Sprintf("%s/%s", c.Directory, groupsFileName)
}

// generateGroupData generates group data and writes it to a CSV file
func (c *Config) generateGroupData() error {
	if c.NumGroups <= 0 {
		return nil
	}

	file, err := os.Create(c.getGroupFilePath())
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
	groups := generateGroupNames(c.NumGroups)

	for _, group := range groups {
		if err := csvWriter.Write([]string{group}); err != nil {
			return err
		}
	}

	// Flush the data to the file
	csvWriter.Flush()

	return nil
}

// generateGroupNames generates a slice of group names using the gofakeit library
// and returns a deduped slice of group names
func generateGroupNames(num int) []string {
	groups := []string{}
	for i := 0; i < num; i++ {
		groups = append(groups, cases.Title(language.English, cases.Compact).String(gofakeit.Adjective()))
	}

	// dedupe the groups
	return sliceutil.Dedupe(groups)
}
