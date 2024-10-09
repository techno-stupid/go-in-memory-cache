package cache

import (
	"time"
)

// StartCleanup starts a background goroutine to clean up expired items
func (c *Cache) StartCleanup(interval time.Duration) {
	go func() {
		for {
			<-time.After(interval)
			c.cleanup()
		}
	}()
}

// cleanup deletes all expired items from the cache
func (c *Cache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now().UnixNano()
	for key, item := range c.items {
		if item.expiration > 0 && now > item.expiration {
			delete(c.items, key)
		}
	}
}
