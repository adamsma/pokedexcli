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
	store map[string]cacheEntry
	mutex *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	newC := Cache{
		mutex: &sync.Mutex{},
		store: make(map[string]cacheEntry),
	}

	go newC.reapLoop(interval)

	return newC
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.store[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.store[key]

	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}

}

func (c *Cache) reap(now time.Time, last time.Duration) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key, val := range c.store {
		if now.Sub(val.createdAt) > last {
			delete(c.store, key)
		}
	}

}
