package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntryMap map[string]cacheEntry
	mu            sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		cacheEntryMap: make(map[string]cacheEntry),
		mu:            sync.Mutex{},
	}
	go newCache.reapLoop(interval)

	return newCache
}

func (c *Cache) reapLoop(interval time.Duration) {
	if interval == 0 {
		return
	}
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.cacheEntryMap {
			if time.Since(v.createdAt).Seconds() > interval.Seconds() {
				delete(c.cacheEntryMap, k)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntryMap[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheEntryMap[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
