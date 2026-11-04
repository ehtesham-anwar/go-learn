package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count uint64
	mu    sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	mu.Lock()
	count++
	defer mu.Unlock()
}

func main() {
	start := time.Now()

	wg := &sync.WaitGroup{}
	for range 10 {
		wg.Add(1)
		go increment(wg)
	}

	wg.Wait()

	fmt.Printf("count value %v\n", count)            // 15.2µs
	fmt.Printf("time taken %v\n", time.Since(start)) // 15.2µs
}
