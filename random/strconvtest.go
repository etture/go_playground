package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "123"
	a_, _ := strconv.ParseInt(a, 10, 64)
	fmt.Println("a_:", a_)
	fmt.Println("a_/1000:", a_/1000)
	fmt.Printf("%d, %d, %f, %f", 1500/1000, 999/1000, 150.0/1000, 999.0/1000)
}
