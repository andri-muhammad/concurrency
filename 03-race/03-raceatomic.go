package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var count int32

	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1)
		}()
	}

	wg.Wait()
	fmt.Println(count)

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}
