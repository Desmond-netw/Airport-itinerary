package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// lookup malformed  var
var (
	ErrMalformedRecord = "malformed itinerary record"
)

// func to lookup in csv file airport codes and return airport name
func lookupInCSV(keyHeader, key, searchHeader, filepath string) (string, bool, error) {
	// open file path and search for key in keyHeader column
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("airport lookup not found")
		os.Exit(1)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	// Read headers
	headers, err := csvReader.Read()
	if err != nil {
		return "", false, fmt.Errorf("%s", ErrMalformedRecord)
	}

	headerIndex := make(map[string]int, len(headers))
	for indx, header := range headers {
		headerIndex[strings.TrimSpace(header)] = indx
	}

	// required headers
	requiredKeyHeader := []string{"name", "iso_country", "municipality", "icao_code", "iata_code", "coordinates"}
	for _, col := range requiredKeyHeader {
		if _, ok := headerIndex[col]; !ok {
			return "", false, fmt.Errorf("%s", ErrMalformedRecord)
		}
	}

	keyIndex, ok := headerIndex[keyHeader]
	searchIndex, ok2 := headerIndex[searchHeader]
	if !ok || !ok2 {
		return "", false, fmt.Errorf("%s", ErrMalformedRecord)
	}

	var result string
	found := false

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", false, fmt.Errorf("%s", ErrMalformedRecord)
		}
		// ensure indices exist in record
		if keyIndex >= len(record) || searchIndex >= len(record) {
			return "", false, fmt.Errorf("%s", ErrMalformedRecord)
		}
		// skip empty key fields
		if strings.TrimSpace(record[keyIndex]) == "" {
			continue
		}
		if strings.EqualFold(strings.TrimSpace(record[keyIndex]), strings.TrimSpace(key)) {
			result = strings.TrimSpace(record[searchIndex])
			found = true
			break
		}
	}

	return result, found, nil
}

// ReadInputFile  reads the content of Itinerary input file
func ReadInputFile(filepath string) (string, bool) {
	// read the file content
	data, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Input file not found:", err)
			os.Exit(1)
		}
		fmt.Println("\nError reading input file:", err)
		return "", false
	}
	return string(data), true
}

// end of file reading

// WriteOutputfile create/ write the  pretttified itinerary to file
func WriteOutputfile(content, filepath string) bool {
	content = RemoveANSI(content)
	err := os.WriteFile(filepath, []byte(content), 0o644)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Output file: not found", err)
			os.Exit(1)
		}
		fmt.Println("Error writing to output file:", err)
		return false
	}
	return err == nil
}
