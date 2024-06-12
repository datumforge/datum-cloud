package seed

import (
	"os"

	"github.com/99designs/gqlgen/graphql"
)

// loadCSVFile loads a CSV file from the file system
func loadCSVFile(fileName string) (graphql.Upload, error) {
	input, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return graphql.Upload{}, err
	}

	return graphql.Upload{
		File:        input,
		Filename:    fileName,
		ContentType: "text/csv",
	}, nil
}
