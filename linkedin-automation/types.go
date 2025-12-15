
package main

// Profile represents a target user from the CSV
type Profile struct {
	ID          string
	Name        string
	LinkedinURL string
	Status      string // e.g., "Success", "Failed"
}

// Config holds the settings
type Config struct {
	LinkedInEmail    string
	LinkedInPassword string
	HeadlessMode     bool // Set to false to see the browser open
	MockMode         bool // Set to true to skip real browser actions (for testing logic)
}