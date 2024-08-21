package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	testKey := "TEST_KEY"
	testValue := "test_value"
	os.Setenv(testKey, testValue)

	result := Config(testKey)
	assert.Equal(t, testValue, result, "Config should return the value set in the environment variable")

	unsetKey := "UNSET_KEY"
	result = Config(unsetKey)
	assert.Equal(t, "", result, "Config should return an empty string for an unset environment variable")

	os.Unsetenv(testKey)
}
