package pokecache

import(
    "time"
)

func (c Cache) Add(key string, val []byte) {
    entry := cacheEntry{
        createdAt: time.Now(),
        val: val,
    }
    c.mu.Lock()
    defer c.mu.Unlock()
    c.cache[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    c.mu.Unlock()
    entry, ok := c.cache[key]
    return entry.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(time.Now().UTC(), interval)
    }
}

func (c Cache) reap(now time.Time, last time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    for key, entry := range c.cache {
        if entry.createdAt.Before(now.Add(-last)) {
            delete(c.cache, key)
        }
    }
}
