package config

import (
	"os"
)

// Config contains global application information
type Config struct {
	Name    string
	Version string
	Commit  string
	Date    string
	WD      string
}

// New creates a new Config from flags and environment variables
func New(version, commit, date string) *Config {
	c := &Config{
		Name:    "Mazerunner",
		Version: version,
		Commit:  commit,
		Date:    date,
	}

	// Get Working Dir
	wd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	c.WD = wd

	return c
}
