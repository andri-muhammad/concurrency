package main

import (
	"fmt"
	"time"
)

// Concurrency is composition of independently executing things (typically, functions).
// A goroutine is a lightweight thread managed by the Go runtime.

func main() {

	start := time.Now()
	go taskA() // Adding the keyword "go" creates a goroutine
	go taskB() // Adding the keyword "go" creates a goroutine

	fmt.Println("\n\nFinished all tasks")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func taskA() {
	time.Sleep(time.Second * 1)
	fmt.Println("\nTask A finished")
}

func taskB() {
	time.Sleep(time.Second * 2)
	fmt.Println("Task B finished")
}
