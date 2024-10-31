package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisClient_Set(t *testing.T) {

	// Mock Redis client
	redisClient := NewRedisClient()

	// Test cases
	tests := []struct {
		key   string
		value interface{}
	}{
		{"key1", "value1"},
		{"key2", 12345},
		{"key3", true},
	}

	for _, tt := range tests {
		err := redisClient.Set(tt.key, tt.value)
		assert.NoError(t, err)

		// Verify the value was set correctly
		val, err := redisClient.Get(tt.key)
		assert.NoError(t, err)
		assert.Equal(t, tt.value, val)
	}
}
