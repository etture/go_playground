package main

import "fmt"

func main() {
	a, b := vals()
	c, d, e := vals2()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	_, f := vals()
	fmt.Println(f)
}

func vals() (int, int) {
	return 3, 7
}

func vals2() (int, int, int) {
	return 3, 7, 10
}
