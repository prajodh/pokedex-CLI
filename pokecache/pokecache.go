package pokecache

import (
	"time"
	"errors"
)

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

type Cache struct{
	cache map[string]cacheEntry
	// time time.Duration
}

func CreateCache(interval time.Duration) Cache{
	Cache_ := Cache{make(map[string]cacheEntry)}
	go Cache_.reaploop(interval)
	return Cache_
}

func (c *Cache) Add(key string, val []byte){
	c.cache[key] = cacheEntry{createdAt: time.Time{}, val : val}
}

func (c * Cache) Get(key string) ([]byte, error){
	val, ok := c.cache[key]
	if ok {
		return val.val, nil
	}
	return val.val, errors.New("empty entry")
}

func (c * Cache) reaploop(interval time.Duration){
	timeChannel := time.NewTicker(interval)
	for range timeChannel.C{
		reap(c, interval)
	}
}

func reap(c *Cache, interval time.Duration){
	time_ := time.Now().UTC().Add(-interval)
	for k, v := range c.cache{
		if v.createdAt.Before(time_){
			delete(c.cache,  k)
		}
	}
}

// func pokecache(){
// 	cache := createCache()
// 	var a []byte
// 	cache.add("hello", a)
// 	fmt.Println(cache)
// 	_, err :=cache.get("b")
// 	fmt.Println(err)

// }