package main

import (
	"fmt"
	"go-for-dev-redis/redis/cache"

	"github.com/rs/zerolog/log"
)

func main() {
	var cache = cache.NewRedisClient()
	var err = cache.Set("key1", "Hello, Redis!")
	if err != nil {
		log.Error().Err(err).Msg("Failed to set value in cache")
		return
	}
	value, err := cache.Get("key1")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get value from cache")
		return
	}
	fmt.Printf("Value set in cache %v", value)

}
