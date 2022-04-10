package helpers

import (
	"encoding/csv"
	"os"
)

func ReadCsv(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, CsvCouldNotOpenFileError
	}

	var closeErr error
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			closeErr = err
		}
	}(f)
	if closeErr != nil {
		return nil, closeErr
	}

	csvReader := csv.NewReader(f)
	csvSlice, err := csvReader.ReadAll()
	if err != nil {
		return nil, CsvFileError
	}

	if len(csvSlice) == 0 {
		return nil, CsvFileEmptyError
	}

	index := -1
	for i, name := range csvSlice[0] {
		if name == "name" {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, CsvColumnsError
	}

	var categoryNames []string
	for _, row := range csvSlice[1:] {
		categoryNames = append(categoryNames, row[index])
	}

	return categoryNames, nil
}
