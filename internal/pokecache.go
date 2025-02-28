package pokecache

import (
	"sync"
	"time"
)

var once sync.Once
var cache *Cache

type Cache struct {
	cached   map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	once.Do(func() {
		cache = &Cache{cached: nil, mutex: sync.Mutex{}, interval: interval}
	})
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	c.cached[key] = cacheEntry{createdAt: time.Now(), val: value}
	defer c.mutex.Unlock()
}

func (c *Cache) Get(key string) (value []byte, in bool) {
	c.mutex.Lock()
	v, in := c.cached[key]
	defer c.mutex.Unlock()
	return v.val, in
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(time.Second * 5)
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for {
		<-ticker.C
		for k, v := range c.cached {
			if time.Now().After(v.createdAt.Add(c.interval)) {
				delete(c.cached, k)
			}
		}
	}
}
