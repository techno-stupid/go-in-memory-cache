package cache

import (
	"sync"
	"time"
)

// CacheItem represents a single cache entry
type CacheItem struct {
	value      interface{}
	expiration int64 // Unix timestamp of expiration time
}

// Cache represents the in-memory cache
type Cache struct {
	items map[string]CacheItem
	mu    sync.RWMutex  // RWMutex allows multiple readers or a single writer at a time
	ttl   time.Duration // Default time-to-live (TTL) for cache items
}
