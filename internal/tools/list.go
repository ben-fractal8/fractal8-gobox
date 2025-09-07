package tools

import (
	"encoding/csv"
	"strings"
)

// AppendOrPrependList modifies each item in the list by adding a prefix or suffix.
func AppendOrPrependList(items []string, prefix, suffix string) []string {
	out := make([]string, len(items))
	for i, item := range items {
		out[i] = prefix + item + suffix
	}
	return out
}

// ListToCSV converts a list of strings into CSV format (one item per row).
func ListToCSV(items []string) string {
	var sb strings.Builder
	writer := csv.NewWriter(&sb)
	writer.Write(items)
	writer.Flush()
	return sb.String()
}

// CSVToList converts CSV data into a list of strings (flattens first column).
func CSVToList(csvData string) ([]string, error) {
	reader := csv.NewReader(strings.NewReader(csvData))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var list []string
	for _, rec := range records {
		if len(rec) > 0 {
			list = append(list, rec[0])
		}
	}
	return list, nil
}
