package files

import (
	"encoding/csv"
	"os"
)

func ReadCSV(dir string) ([][]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
