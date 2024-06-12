package seed

import (
	"encoding/csv"
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

// getColumnIndex returns the index of a column in a CSV file
// if the column does not exist, it returns -1
func getColumnIndex(headers []string, columnName string) int {
	for i, header := range headers {
		if header == columnName {
			return i
		}
	}

	return -1
}

// readCSVFile reads a CSV file from the file system
func readCSVFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	return reader.ReadAll()
}
