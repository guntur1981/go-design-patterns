package main

import (
	"fmt"
	"sync"
)

type Database interface {
	GetValue(key string) int
}

type CacheDB struct {
	cache map[string]int
}

func (c *CacheDB) init() {
	if c.cache == nil {
		c.cache = map[string]int{}
	}
	c.cache["foo"] = 1
	c.cache["bar"] = 2
	c.cache["baz"] = 3
	// populate a lot of data
	// ...
}

func (c CacheDB) GetValue(key string) int {
	return c.cache[key]
}

var (
	instanceDB *CacheDB
	once       sync.Once
)

func NewDB() *CacheDB {
	once.Do(func() {
		instanceDB = &CacheDB{}
		instanceDB.init()
	})
	return instanceDB
}

func Sum(db Database, keys []string) int {
	sum := 0
	for _, key := range keys {
		sum += db.GetValue(key)
	}
	return sum
}

func main() {
	db := NewDB() // always get the singleton reference
	total := Sum(db, []string{"foo", "baz"})
	fmt.Println("Total:", total)
}
