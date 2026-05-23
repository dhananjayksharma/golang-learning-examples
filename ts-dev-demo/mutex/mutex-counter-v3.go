package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu    sync.RWMutex
	store map[string]string
}

func (c *Cache) Get(k string) (string, bool) {
	c.mu.RLock() // multiple readers can hold this
	defer c.mu.RUnlock()
	v, ok := c.store[k]
	return v, ok
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock() // exclusive write
	defer c.mu.Unlock()
	c.store[k] = v
}

func main() {
	c := &Cache{
		store: make(map[string]string),
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Set(fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(c.Get(fmt.Sprintf("key-%d", i)))
		}(i)
	}

	wg.Wait()

	if v, ok := c.Get("key-42"); ok {
		fmt.Println(v)
	}
}
