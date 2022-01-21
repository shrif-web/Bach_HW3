package main

import (
	"os"
)

func main() {
	addr := os.Getenv("CACHE_ADDR")
	if addr == "" {
		addr = "localhost:3124"
	}
	sv := NewServer()

	sv.Run("3000")
}
