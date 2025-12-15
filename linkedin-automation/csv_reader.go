
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadLeads reads the CSV file and returns a slice of Profiles
func ReadLeads(filename string) ([]Profile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not parse CSV: %w", err)
	}

	var profiles []Profile
	// Start from i=1 to skip the Header row
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 3 {
			continue // Skip invalid rows
		}
		profiles = append(profiles, Profile{
			ID:          row[0],
			Name:        row[1],
			LinkedinURL: row[2],
			Status:      "Pending",
		})
	}
	return profiles, nil
}