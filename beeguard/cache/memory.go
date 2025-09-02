package beeguard

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
)

type MemoryCache struct {
	sync.Mutex
	cache map[string]string
}

func NewMemoryCache() *MemoryCache {
	slog.Debug("Default bee-memory cache initialized")
	return &MemoryCache{cache: make(map[string]string)}
}

func (c *MemoryCache) Set(ctx context.Context, key string, value string) (err error) {
	c.Lock()
	c.cache[key] = value
	c.Unlock()

	fmt.Println("====== Memory cache updated =======")
	for key, val := range c.cache {
		fmt.Printf("key %s val %s\n", key, val)
	}
	fmt.Println("=============")
	return
}

func (c *MemoryCache) Get(ctx context.Context, key string) (string, error) {
	c.Lock()
	defer c.Unlock()
	val, ok := c.cache[key]
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}
	return val, nil
}
