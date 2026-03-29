package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			m.Store("key", key)
		}(i)
	}

	wg.Wait()

	value, _ := m.Load("key")
	fmt.Printf("Value: %v\n", value)
}
