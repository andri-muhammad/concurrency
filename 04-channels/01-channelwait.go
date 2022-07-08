package main

import (
	"fmt"
	"time"
)

// Channels are used to pass data between/synchronize goroutines, including func main
// Do not communicate by sharing memory; instead, share memory by communicating

func main() {

	ch := make(chan string)

	go sendMe(ch)
	fmt.Println(<-ch)

}

func sendMe(ch chan<- string) {

	time.Sleep(time.Second * 2)
	ch <- "SendMe is done"

}
