package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a1 := a[2:7]
	fmt.Println("a:", a)
	fmt.Println("a1:", a1)
	a2 := a[4:]
	fmt.Println("a2:", a2)
	a3 := a2[:3]
	fmt.Println("a3:", a3)
	a4 := append(a3, 1000, 1001)
	fmt.Println("a4:", a4)
	fmt.Println("a3:", a3)
	fmt.Println("a:", a)
	fmt.Println("a2:", a2)

	b := []byte("hello ")
	fmt.Println("b:", b)
	b1 := append(b, "world"...)
	fmt.Println("b:", b)
	fmt.Println("b1:", b1)
}
