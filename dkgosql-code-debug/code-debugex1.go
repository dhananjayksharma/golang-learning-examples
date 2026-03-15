package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://service-a/api",
		"https://service-b/api",
		"https://service-a/api",
	}

	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: mockTransport{},
	}
	cache := NewCache()

	var wg sync.WaitGroup
	for _, u := range urls {
		url := u
		wg.Add(1)
		go func() {
			defer wg.Done()
			val, err := fetch(client, cache, url)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
			fmt.Println(val)
		}()
	}
	wg.Wait()
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func fetch(client *http.Client, cache *Cache, url string) (string, error) {
	if val, ok := cache.Get(url); ok {
		return val, nil
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	val := string(body)
	cache.Set(url, val)
	return val, nil
}

type mockTransport struct{}

func (mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	responses := map[string]string{
		"https://service-a/api": "service-a: cached response",
		"https://service-b/api": "service-b: fresh response",
	}

	body, ok := responses[req.URL.String()]
	if !ok {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Body:       io.NopCloser(strings.NewReader("not found")),
			Header:     make(http.Header),
			Request:    req,
		}, nil
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
		Request:    req,
	}, nil
}
