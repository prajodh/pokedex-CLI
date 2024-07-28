package pokecache

import (
	"fmt"
	"time"
)

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func Pokecache(){
	fmt.Println("hello")
}