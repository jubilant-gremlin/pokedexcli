package pokecache

import (
    "sync"
    "time"
) 
type Cache struct {
    cache map[string]cacheEntry
    mu *sync.Mutex
}

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

func NewCache(interval time.Duration) Cache{
    new := Cache{
        cache: make(map[string]cacheEntry),
        mu: &sync.Mutex{},
    }

    go new.reapLoop(interval)

    return new
}
