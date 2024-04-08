package exporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigValidation_NoError(t *testing.T) {
	cfg := &config{}

	err := cfg.Validate()

	assert.NoError(t, err)
}
