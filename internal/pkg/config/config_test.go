package config

import (
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

// I used to think my life was a tragedy. But now I realise it's a fcking comedy

// TestConfig tests config
func TestConfig(t *testing.T) {
	cfgCopy := GetConfig()
	cfgExpected := Config{}

	b, err := os.ReadFile(cfgFilename)
	if err != nil {
		t.Error(err)
	}

	err = yaml.Unmarshal(b, &cfgExpected)
	if err != nil {
		t.Error(err)
	}

	if cfgExpected != cfgCopy {
		t.Errorf("Expected:\n\t%+v, but got:\n\t%+v", cfgExpected, cfgCopy)
	}
}
