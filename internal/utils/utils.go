package utils

import (
	"fmt"
	"os"
)

var osStat = os.Stat

// Fonction pour afficher un message
func PrintHello() {
	fmt.Println("Hello from utils package")
}

// IsRunningInDocker checks if the application is running inside a Docker container
func IsRunningInDocker() bool {
	_, err := osStat("/.dockerenv")
	return err == nil
}
