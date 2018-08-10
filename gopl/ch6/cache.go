// version 1.0 2018-08-10
package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func Save(key, value string) {
	cache.Lock()
	cache.mapping[key] = value
	cache.Unlock()
}

func main() {
	key, value := "key", "value"
	Save(key, value)
	fmt.Println(Lookup(key))
}
