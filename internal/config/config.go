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
	Env     string
}

// New creates a new Config from flags and environment variables
func New(version, commit, date string, local bool) *Config {
	var env string
	if local {
		env = "local"
	} else {
		env = "production"
	}

	c := &Config{
		Name:    "Mazerunner",
		Version: version,
		Commit:  commit,
		Date:    date,
		Env:     env,
	}

	// Get Working Dir
	wd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	c.WD = wd

	return c
}
