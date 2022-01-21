package main

import (
	"fmt"
	"os"

	"server/cache"
)

func main() {
	addr := os.Getenv("CACHE_ADDR")
	if addr == "" {
		addr = "localhost:3124"
	}
	fmt.Println("here")
	client := cache.NewCacheClient(addr)
	client.Put("asd", "asdd")
	fmt.Println(client.Get("asdf"))
	println(client.Get("asd"))
	client.Put("asdf", "asddff")
	println(client.Get("asdf"))
}
