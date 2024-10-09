package main

import (
	"fmt"
	"github.com/techno-stupid/go-in-memory-cache"
	"time"
)

func main() {
	// Create a new cache with a default TTL of 5 seconds
	c := cache.New(5 * time.Second)
	// Set an item in the cache with unlimited TTL
	c.Set("marco", "polo", 0)
	// Set an item in the cache
	c.Set("foo", "bar", 10*time.Second)
	time.Sleep(6 * time.Second)

	// Retrieve the item
	value, found := c.Get("foo")
	if found {
		fmt.Println("Found value:", value)
	}

	// Wait for 6 seconds (item should expire)
	time.Sleep(6 * time.Second)

	// Try to get the expired item
	_, found = c.Get("foo")
	if !found {
		fmt.Println("Item expired and not found")
	}
	value, found = c.Get("marco")
	if found {
		fmt.Println("Found value:", value)
	}
}
