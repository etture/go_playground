package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 5000000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Printf("%T\n", n)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	const a float64 = 4.22
	fmt.Println(a)
}
