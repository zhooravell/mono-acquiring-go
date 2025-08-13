package monoacquiring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Validate_EmptyAPIKey(t *testing.T) {
	config := Config{
		APIKey: "",
	}
	err := config.Validate()

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrBlankAPIKey)
	assert.Equal(t, "api key is blank", err.Error())
}

func TestConfig_Validate_EmptyBaseURL(t *testing.T) {
	config := Config{
		APIKey:  "test",
		BaseURL: "",
	}
	err := config.Validate()

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrBlankBaseURL)
	assert.Equal(t, "base URL is blank", err.Error())
}

func TestConfig_Validate_InvalidBaseURL(t *testing.T) {
	config := Config{
		APIKey:  "test",
		BaseURL: "http://example.com:port",
	}
	err := config.Validate()

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidBaseURL)
}
