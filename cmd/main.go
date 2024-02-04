package main

import (
	"fmt"
	"github.com/ArmiKGG/in_memory_cache/internal/cache"
)

func main() {
	manage := cache.New()

	s := manage.Set("Arman", map[string]any{"a": 1})
	fmt.Println(s)
	val, ok := manage.Get("Arman")
	fmt.Println(val, ok)
	d := manage.Delete("Arman")
	fmt.Println(d)
	val, ok = manage.Get("Arman")
	fmt.Println(val, ok)
}
