package seed

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

const (
	subscribersFileName = "subscribers.csv"
)

// getSubscriberFilePath returns the full path to the subscribers file
func (c *Config) getSubscriberFilePath() string {
	return fmt.Sprintf("%s/%s", c.Directory, subscribersFileName)
}

// generateSubscriberData generates subscriber data and writes it to a CSV file
func (c *Config) generateSubscriberData() error {
	if c.NumSubscribers <= 0 {
		return nil
	}

	file, err := os.Create(c.getSubscriberFilePath())
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	// Add column headers
	if err := csvWriter.Write([]string{"Email"}); err != nil {
		return err
	}

	// Add data
	for range c.NumSubscribers {
		p := gofakeit.Person()

		if err := csvWriter.Write([]string{fmt.Sprintf("%s.%s@example.com", strings.ToLower(p.FirstName), strings.ToLower(p.LastName))}); err != nil {
			return err
		}
	}

	// Flush the data to the file
	csvWriter.Flush()

	return nil
}
