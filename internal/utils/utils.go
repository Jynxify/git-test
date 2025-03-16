package utils

import (
	"fmt"
	"os"
)

var osStat = os.Stat

// IsRunningInDocker checks if the application is running inside a Docker container
func IsRunningInDocker() bool {
	fmt.Println("Checking if running in Docker")
	_, err := osStat("/.dockerenv")
	return err == nil
}