package cache

import (
	"time"
)

// New initializes a new in-memory cache with a default TTL
func New(defaultTTL time.Duration) *Cache {
	c := &Cache{
		items: make(map[string]CacheItem),
		ttl:   defaultTTL,
	}
	go c.StartCleanup(10 * time.Second)
	return c
}

// Set adds a value to the cache with an optional custom TTL. If TTL is 0, the item never expires.
func (c *Cache) Set(key string, value interface{}, customTTL ...time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ttl := c.ttl
	if len(customTTL) > 0 {
		ttl = customTTL[0]
	}

	var expiration int64
	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	} else {
		expiration = 0 // Indicates the item should never expire
	}

	c.items[key] = CacheItem{
		value:      value,
		expiration: expiration,
	}
}

// Get retrieves a value from the cache if it has not expired
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.items[key]
	if !found {
		c.mu.RUnlock()
		return nil, false
	}

	// Check if the item has expired, unless expiration is set to 0 (unlimited time)
	if item.expiration > 0 && time.Now().UnixNano() > item.expiration {
		c.mu.RUnlock()

		// Now acquire a write lock to delete the expired item
		c.Delete(key)
		return nil, false
	}

	c.mu.RUnlock()
	return item.value, true
}

// Delete removes a key from the cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}
