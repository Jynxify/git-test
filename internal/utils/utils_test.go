package utils

import (
	"errors"
	"os"
	"testing"
)

func TestIsRunningInDocker(t *testing.T) {
	originalStat := osStat
	defer func() { osStat = originalStat }()

	tests := []struct {
		name     string
		mockStat func(string) (os.FileInfo, error)
		expected bool
	}{
		{
			name: "Not running in Docker",
			mockStat: func(string) (os.FileInfo, error) {
				return nil, errors.New("file not found")
			},
			expected: false,
		},
		{
			name: "Running in Docker",
			mockStat: func(string) (os.FileInfo, error) {
				return nil, nil
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			osStat = tt.mockStat
			result := IsRunningInDocker()
			if result != tt.expected {
				t.Errorf("IsRunningInDocker() = %v; want %v", result, tt.expected)
			}
		})
	}
}
