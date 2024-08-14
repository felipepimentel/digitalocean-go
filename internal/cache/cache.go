package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/felipepimentel/digitalocean-go/internal/logging"
)

type Cache struct {
	Path string
}

type CacheEntry struct {
	Data      interface{}
	ExpiresAt time.Time
}

func New(path string) *Cache {
	return &Cache{Path: path}
}

func (c *Cache) Set(key string, data interface{}, expiration time.Duration) error {
	entry := CacheEntry{
		Data:      data,
		ExpiresAt: time.Now().Add(expiration),
	}

	bytes, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	filename := filepath.Join(c.Path, key)
	return os.WriteFile(filename, bytes, 0644)
}

func (c *Cache) Get(key string, result interface{}) bool {
	filename := filepath.Join(c.Path, key)
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return false
	}

	var entry CacheEntry
	if err := json.Unmarshal(bytes, &entry); err != nil {
		logging.ErrorLogger.Printf("Failed to unmarshal cache entry: %v", err)
		return false
	}

	if time.Now().After(entry.ExpiresAt) {
		os.Remove(filename)
		return false
	}

	if err := json.Unmarshal(bytes, result); err != nil {
		logging.ErrorLogger.Printf("Failed to unmarshal cached data: %v", err)
		return false
	}

	return true
}