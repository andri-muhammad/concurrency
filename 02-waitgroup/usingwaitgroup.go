package main

import (
	"fmt"
	"sync"
	"time"
)

// A WaitGroup waits for a collection of goroutines to finish.
var wg = sync.WaitGroup{} // This creates our waitgroup called, "wg"

func main() {

	wg.Add(2) // This sets our wait group to 2, meaning we will wait for 2 goroutines to complete
	start := time.Now()
	go taskA()
	go taskB()

	wg.Wait() // This blocks (waits) until wg=0 meaning all goroutines are done.
	fmt.Println("\n\nFinished all tasks")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func taskA() {
	time.Sleep(time.Second * 1)
	fmt.Println("\nTask A finished")
	wg.Done() // This decrements wg by one, indicating that taskA is done.
}

func taskB() {
	time.Sleep(time.Second * 2)
	fmt.Println("Task B finished")
	wg.Done() // This decrements wg by one, indicating that taskB is done.
}
