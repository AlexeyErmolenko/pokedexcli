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
	data            map[string]cacheEntry
	rwMu            *sync.RWMutex
	defaultDuration time.Duration
}

func NewCache(defaultDuration time.Duration) *Cache {
	cache := Cache{
		data:            make(map[string]cacheEntry),
		rwMu:            &sync.RWMutex{},
		defaultDuration: defaultDuration,
	}

	go cache.reapLoop()

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	cache := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.data[key] = cache

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	value, ok := c.data[key]

	if !ok {
		return nil, false
	}

	return value.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.defaultDuration * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		c.reap(t)
	}
}

func (c *Cache) reap(t time.Time) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()

	for i, v := range c.data {
		if t.Sub(v.createdAt).Seconds() > c.defaultDuration.Seconds() {
			delete(c.data, i)
		}
	}
}
