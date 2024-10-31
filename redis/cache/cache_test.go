package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// MockCache is a mock implementation of the ICache interface for testing purposes.
type MockCache struct {
	store map[string]interface{}
}

func NewMockCache() *MockCache {
	return &MockCache{store: make(map[string]interface{})}
}

func (m *MockCache) Set(key string, value interface{}) error {
	m.store[key] = value
	return nil
}

func (m *MockCache) SetWithDuration(key string, value interface{}, duration int) error {
	m.store[key] = value
	go func() {
		time.Sleep(time.Duration(duration) * time.Second)
		delete(m.store, key)
	}()
	return nil
}

func (m *MockCache) Get(key string) (interface{}, error) {
	value, exists := m.store[key]
	if !exists {
		return nil, nil
	}
	return value, nil
}

func (m *MockCache) Delete(key string) error {
	delete(m.store, key)
	return nil
}

func (m *MockCache) Update(key string, value interface{}) error {
	if _, exists := m.store[key]; !exists {
		return nil
	}
	m.store[key] = value
	return nil
}

func TestSet(t *testing.T) {
	cache := NewMockCache()
	err := cache.Set("key1", "value1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", cache.store["key1"])
}

func TestSetWithDuration(t *testing.T) {
	cache := NewMockCache()
	err := cache.SetWithDuration("key1", "value1", 1)
	assert.NoError(t, err)
	assert.Equal(t, "value1", cache.store["key1"])
	time.Sleep(2 * time.Second)
	_, exists := cache.store["key1"]
	assert.False(t, exists)
}

func TestGet(t *testing.T) {
	cache := NewMockCache()
	cache.Set("key1", "value1")
	value, err := cache.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", value)
}

func TestDelete(t *testing.T) {
	cache := NewMockCache()
	cache.Set("key1", "value1")
	err := cache.Delete("key1")
	assert.NoError(t, err)
	_, exists := cache.store["key1"]
	assert.False(t, exists)
}

func TestUpdate(t *testing.T) {
	cache := NewMockCache()
	cache.Set("key1", "value1")
	err := cache.Update("key1", "value2")
	assert.NoError(t, err)
	assert.Equal(t, "value2", cache.store["key1"])
}
