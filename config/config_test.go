package config

import (
	"fmt"
	"os"
	"testing"
)

func TestNew_PanicsWhenEnvConfigFileNotFound(t *testing.T) {
	defer func() {
		osEnvFunction = os.Getenv
		fmt.Println("osEnvFunction reset back to os.Getenv")
	}()
	defer func() {
		if r := recover(); r == nil {
			t.Error(t, "The New function should have panicked")
		}
	}()
	osEnvFunction = func(string) string {
		fmt.Println("in the mock")
		return "Not an Environment. Should panic"
	}

	// The following is the code under test
	New("./../environments")
}

func TestGetStringSetting(t *testing.T) {
	config := New("./../environments")

	// The logs.level setting should always be at debug for local testing
	got := config.GetStringSetting("logs.level")
	if got != "debug" {
		fmt.Printf("Expected debug got %v", got)
	}
}
