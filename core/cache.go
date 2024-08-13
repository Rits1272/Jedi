package core 

import (
  "sync"
)

/*
Metadata - piggybacking the counter in the metadata along with timestamp

--------------------------------------
counter        |        timestamp
--------------------------------------
8 bit          |        24 bit
--------------------------------------
*/
type CacheItem struct {
  Metadata     uint32
  Value        interface{}
}

type Cache struct {
  items map[interface{}]*CacheItem
  mu    sync.RWMutex
}

func NewCache() *Cache {
  return &Cache {
    items: make(map[interface{}]*CacheItem),
  }
}

func (c *Cache) WithLock(f func()) {
  c.mu.Lock()
  defer c.mu.Unlock()
  f()
}

func (c *Cache) WithRLock(f func()) {
  c.mu.RLock()
  defer c.mu.RUnlock()
  f()
}

func (c *Cache) SetKey(k string, v interface{}) {
  c.WithLock(func() {
    c.items[k] = &CacheItem {
      Metadata: GetCurrentTimestamp(),
      Value: v,
    }
  })
}

func (c *Cache) GetKey(k string) *CacheItem {
  var item *CacheItem

  c.WithLock(func() {
    ci, exists := c.items[k]

    if !exists {
      item = nil
    } else {
      ci.updateMetadata()
      item = ci
    }
  })

  return item
}

func (c *Cache) DelKey(k string) *CacheItem {
  var item *CacheItem

  c.WithLock(func() {
    ci, exists := c.items[k]

    if !exists {
      item = nil
    } else {
      item = ci
      delete(c.items, k)
    }
  })

  return item
}

// Note: Should be used inside Lock
func (ci *CacheItem) updateMetadata() {
  timestamp := GetCurrentTimestamp()
  counter := GetLogCounter(ci.Metadata)

  ci.Metadata = (uint32(counter) << 24) | timestamp
}

var Jedi = NewCache()
