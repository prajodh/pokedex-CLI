package pokecache

import (
	"fmt"
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

func createCache() Cache{
	cache := Cache{make(map[string]cacheEntry)}
	return cache
}

func (c *Cache) add(key string, val []byte){
	c.cache[key] = cacheEntry{createdAt: time.Time{}, val : val}
}

func (c * Cache) get(key string) (cacheEntry, error){
	val, ok := c.cache[key]
	if ok {
		return val, nil
	}
	return val, errors.New("empty entry")
}

func Pokecache(){
	cache := createCache()
	var a []byte
	cache.add("hello", a)
	fmt.Println(cache)
	_, err :=cache.get("b")
	fmt.Println(err)

}