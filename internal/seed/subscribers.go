package seed

import (
	"encoding/csv"
	"fmt"
	"os"

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
		if err := csvWriter.Write([]string{gofakeit.Email()}); err != nil {
			return err
		}
	}

	// Flush the data to the file
	csvWriter.Flush()

	return nil
}
