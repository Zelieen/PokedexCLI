package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]cacheEntry
	mu   sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{}
	c.reapLoop(interval)
	return c
}

func (c Cache) Add(key string, rawBytes []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       rawBytes,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, ok
	}
	val := entry.val
	return val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
	//should use time.Ticker to start reaping after interval
	//reaping: clear all cache entries that are older than the interval
}
