package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func main() {

	start := time.Now()
	go taskA()
	go taskB()

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("\n\nFinished all tasks")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func taskA() {
	time.Sleep(time.Second * 1)
	ch <- "task A finished"
}

func taskB() {
	time.Sleep(time.Second * 2)
	ch <- "Task B finished"
}
