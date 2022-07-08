package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}

	wg.Wait()
	fmt.Println(count)

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}
