package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	queue2 := make(chan int, 3)
	done := make(chan bool)

	go func() {
		for e := range queue2 {
			fmt.Println("received", e)
		}
		done <- true
	}()

	for i := 0; i < 5; i++ {
		queue2 <- i
		time.Sleep(1 * time.Second)
	}
	close(queue2)

	<-done
}
