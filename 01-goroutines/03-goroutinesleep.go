package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	go taskA()
	go taskB()

	time.Sleep(time.Second * 2) // Don't do this in production.  VERY inefficient and unpredictable.

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
