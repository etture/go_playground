package main

import (
	"fmt"
	"runtime"
)

func main() {
	whileLoop()
	chanSelect()
	waitForGoroutine()
}

func whileLoop() {
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 10 {
			break
		}
	}
	fmt.Println()
}

func chanSelect() {
	c, quit := make(chan int), make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
	fmt.Println()
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func waitForGoroutine() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Running:", i)
			if i == 2 {
				c <- 1
			}
		}
		c <- 1
	}()

	<-c

	//time.Sleep(1 * time.Second)
	runtime.GC()
}
