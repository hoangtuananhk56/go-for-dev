package cache

// ICache defines the interface for a cache system.
// It provides methods to set, get, delete, and update cache entries.
type ICache interface {
	// Set stores the given value associated with the key in the cache.
	// Returns an error if the operation fails.
	Set(key string, value interface{}) error

	// SetWithDuration stores the given value associated with the key in the cache
	// with a specified duration (in seconds) after which the entry expires.
	// Returns an error if the operation fails.
	SetWithDuration(key string, value interface{}, duration int) error

	// Get retrieves the value associated with the given key from the cache.
	// Returns the value and an error if the operation fails or the key does not exist.
	Get(key string) (interface{}, error)

	// Delete removes the entry associated with the given key from the cache.
	// Returns an error if the operation fails.
	Delete(key string) error

	// Update modifies the value associated with the given key in the cache.
	// Returns an error if the operation fails.
	Update(key string, value interface{}) error
}
