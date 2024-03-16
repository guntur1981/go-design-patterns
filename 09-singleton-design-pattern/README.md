# Description

The Singleton Design Pattern allows you to ensure that a struct only has one instance.

## Example

Consider the following example:

```
package main

import "fmt"

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

func NewDB() *CacheDB {
	db := CacheDB{}
	db.init()
	return &db
}

func Sum(db *CacheDB, keys []string) int {
	sum := 0
	for _, key := range keys {
		sum += db.GetValue(key)
	}
	return sum
}

func main() {
	db := NewDB()
	total := Sum(db, []string{"foo", "baz"})
	fmt.Println("Total:", total)
}

```

In the above example, we simulated a CacheDB which requires a lot of data to initialize.

## The Reason Why This Is Not a Good Practice

1. If the construction process is costly, this is one reason to use a singleton.
2. This violates the Dependency Inversion Principle where a low-level module (CacheDB) is accessed directly by function Sum().

## A Better Approach

**First**, to adhere to the Dependency Inversion Principle, let's create an interface to provide an abstraction to low-level module.

```
type Database interface {
	GetValue(key string) int
}
```

**Second**, create a package level variable that acts as a singleton that stores a pointer to a database reference.

```
var instanceDB *CacheDB
```

**Third**, we use `sync.Once` to ensure the construction only happens once. Whenever the NewDB() function is called, it always returns the singleton.

```
var once sync.Once

func NewDB() *CacheDB {
	once.Do(func() {
		instanceDB = &CacheDB{}
		instanceDB.init()
	})
	return instanceDB
}
```

**Finally**, we modify the Sum() function using the Database interface.

```
func Sum(db Database, keys []string) int {
	sum := 0
	for _, key := range keys {
		sum += db.GetValue(key)
	}
	return sum
}
```
