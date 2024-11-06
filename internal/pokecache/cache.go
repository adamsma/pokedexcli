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
	store    map[string]cacheEntry
	mutex    *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	var newC Cache
	newC.interval = interval
	newC.mutex = &sync.Mutex{}
	newC.store = map[string]cacheEntry{}

	ticker := time.NewTicker(interval)

	go func() {
		for {
			t := <-ticker.C
			newC.reapLoop(t)
		}
	}()

	return newC
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.store[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {

	entry, ok := c.store[key]

	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(t time.Time) {

	for key, entry := range c.store {
		if t.Sub(entry.createdAt) > c.interval {
			c.mutex.Lock()
			delete(c.store, key)
			c.mutex.Unlock()
		}
	}

}
