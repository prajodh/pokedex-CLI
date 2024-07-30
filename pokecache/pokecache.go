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

func CreateCache() Cache{
	Cache_ := Cache{make(map[string]cacheEntry)}
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

// func pokecache(){
// 	cache := createCache()
// 	var a []byte
// 	cache.add("hello", a)
// 	fmt.Println(cache)
// 	_, err :=cache.get("b")
// 	fmt.Println(err)

// }