package internal

import (
	"testing"
)

func TestConfiguracion(t *testing.T) {
	tempConfigFile := "../test_config.yaml"

	config, err := loadConfig(tempConfigFile)
	if err != nil {
		t.Fatalf("Error loading config: %v", err)
	}

	if config == nil {
		t.Error("Expected non-nil config, got nil")
	}
}