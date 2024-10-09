package cache

import (
	"testing"
	"time"
)

func TestCacheSetGet(t *testing.T) {
	c := New(1 * time.Second)
	c.Set("foo", "bar")

	value, found := c.Get("foo")
	if !found || value != "bar" {
		t.Errorf("expected to find 'bar', found %v", value)
	}
}

func TestCacheExpiration(t *testing.T) {
	c := New(1 * time.Second)
	c.Set("foo", "bar")

	time.Sleep(2 * time.Second)

	_, found := c.Get("foo")
	if found {
		t.Errorf("expected 'foo' to expire, but it was found")
	}
}

func TestCacheCustomTTL(t *testing.T) {
	c := New(5 * time.Second)          // Default TTL is 5 seconds
	c.Set("foo", "bar", 1*time.Second) // Custom TTL for this key

	time.Sleep(2 * time.Second)

	_, found := c.Get("foo")
	if found {
		t.Errorf("expected 'foo' to expire with custom TTL")
	}
}
