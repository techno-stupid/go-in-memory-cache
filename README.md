# Go In-Memory Cache Package

A lightweight, thread-safe, in-memory cache for Go with support for TTL (Time-to-Live) per item, automatic cleanup, and items with unlimited expiration time.

## Features

- Set and get cache items with a configurable TTL (Time-to-Live).
- Support for items that never expire (TTL of 0).
- Thread-safe with concurrent access control using sync.RWMutex.
- Background cleanup process to remove expired items automatically.
- Custom TTL for individual items or use a global default TTL.
- Manual item deletion.

## Installation

```bash
go get github.com/techno-stupid/go-in-memory-cache
```

## Usage

### 1. Create a Cache with Default TTL
You can create a cache with a default TTL (e.g., 5 seconds). All items added to the cache will expire after the default TTL unless a custom TTL is specified.

```go
package main

import (
	"fmt"
	"time"
	"github.com/techno-stupid/go-in-memory-cache"
)

func main() {
	// Create a cache with a default TTL of 5 seconds
	c := cache.New(5 * time.Second)

	// Add an item to the cache with the default TTL
	c.Set("foo", "bar")

	// Retrieve the item before it expires
	value, found := c.Get("foo")
	if found {
		fmt.Println("Found value:", value)
	} else {
		fmt.Println("Value expired or not found")
	}
}
```

### 2. Add Items with Custom TTL
You can specify a custom TTL for each item when calling the Set method. For example, adding an item with a TTL of 10 seconds:

```go
// Set an item with a custom TTL of 10 seconds
c.Set("baz", "qux", 10*time.Second)
```

### 3. Set Items with No Expiration (Unlimited Time)
To add an item that never expires, pass a TTL of 0 when calling Set:

```go
// Set an item that will never expire (TTL = 0)
c.Set("foo", "bar", 0)

// This item can be accessed indefinitely until manually deleted
value, found := c.Get("foo")
if found {
    fmt.Println("Found value:", value)
}
```
### 4. Manually Deleting Items
You can create a cache with a default TTL (e.g., 5 seconds). All items added to the cache will expire after the default TTL unless a custom TTL is specified.

```go
// Manually delete an item
c.Delete("foo")
```

## Examples

```go
package main

import (
    "fmt"
    "time"
	"github.com/techno-stupid/go-in-memory-cache"
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
```

## Thread Safety
This package is designed to be **thread-safe**. It uses `sync.RWMutex` to handle concurrent read and write operations to the cache, ensuring safe access across multiple goroutines.

## Summary
- The package allows you to create a cache with customizable TTL, background cleanup, and support for items with unlimited expiration.
- It is lightweight, simple to use, and thread-safe.


Let me know if you need further customization!






