package utils

import (
	"os"
)

var osStat = os.Stat

// IsRunningInDocker checks if the application is running inside a Docker container
func IsRunningInDocker() bool {
	_, err := osStat("/.dockerenv")
	return err == nil
}