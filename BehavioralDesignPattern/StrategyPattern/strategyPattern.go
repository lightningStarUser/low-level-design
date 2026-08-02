// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Cache struct {
	storage     map[int]string
	evicAlgo    EvictionAlgo
	capacity    int
	maxCapacity int
}

func (c *Cache) setEvictionAlgo(evicAlgo EvictionAlgo) {
	c.evicAlgo = evicAlgo
}

func (c *Cache) add(key int, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.storage[key] = value
	c.capacity++
}

func (c *Cache) evict() {
	c.evicAlgo.evict(c)
	c.capacity--
}

type EvictionAlgo interface {
	evict(c *Cache)
}

type LRU struct{}
type FIFO struct{}
type LFU struct{}

func (l *LRU) evict(c *Cache) {
	fmt.Println("Eviction by LRU strategy")
}

func (l *FIFO) evict(c *Cache) {
	fmt.Println("Eviction by FIFO strategy")
}

func (l *LFU) evict(c *Cache) {
	fmt.Println("Eviction by LFU strategy")
}

func initCache(e EvictionAlgo) *Cache {
	return &Cache{
		storage:     make(map[int]string),
		evicAlgo:    e,
		capacity:    0,
		maxCapacity: 2,
	}
}

func main() {
	lru := &LRU{}
	cache := initCache(lru)
	cache.add(1, "Anna")
	cache.add(2, "Bella")
	cache.add(3, "Celia")
	lfu := &LFU{}
	cache = initCache(lfu)
	cache.add(1, "Anna")
	cache.add(2, "Bella")
	cache.add(3, "Celia")
	fifo := &FIFO{}
	cache = initCache(fifo)
	cache.add(1, "Anna")
	cache.add(2, "Bella")
	cache.add(3, "Celia")
}
