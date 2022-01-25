package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hi")
	ifTest(10)
	//err := unnecessaryElse("yo")
	//if err != nil {
	//	return
	//}
	testLabeledBlock()
	testMap()
	testRange()
	testNewKeyword()
	testCompositeLiteral()
	testNewVsMake()
	test2DSlices()
	testFindMin()
	testAppend()
	testConst()
}

func ifTest(num int) {
	if x := num * 2; x > 10 {
		fmt.Println("over 10!")
	} else {
		fmt.Println("not over 10...")
	}
}

func unnecessaryElse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	fmt.Println(f)
	return err
}

func testLabeledBlock() {
Loop:
	for n := 0; n < 10; n++ {
		fmt.Println(n)
		if n >= 5 {
			break Loop
		}
	}
}

func testForLoop1() {
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			fmt.Printf("num: %v\n", i)
		}
	}
}

func testForLoop2() {

}

func testMap() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	val, prs := m["k2"]
	fmt.Println("val:", val)
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	fmt.Println()
}

func testRange() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index when 3:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "some string" {
		fmt.Println(i, string(c))
	}

	fmt.Println()

}

type MyType struct {
	a int
	b string
}

func testNewKeyword() {
	myType := new(MyType)
	fmt.Println(myType)
	fmt.Println(myType.a)
	fmt.Println()
}

func testCompositeLiteral() {
	const Enone, Eio, Einval = 1, 2, 3
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	fmt.Println(a, s, m)
	fmt.Println()
}

func testNewVsMake() {
	var p *[]int = new([]int)
	var v []int = make([]int, 100)
	fmt.Printf("p: %T, v: %T\n", p, v)

	var pp *[]int = new([]int)
	*pp = make([]int, 100, 100)

	vv := make([]int, 100)
	fmt.Println(vv)
	fmt.Println()
}

func test2DSlices() {
	type LinesOfText [][]byte

	YSize, XSize := 5, 4

	// allocate each line
	picture := make([][]uint8, YSize)
	for i := range picture {
		picture[i] = make([]uint8, XSize)
	}
	fmt.Println(picture)

	// allocate single slice
	picture2 := make([][]uint8, YSize)
	pixels := make([]uint8, XSize*YSize)
	for i := range picture2 {
		picture2[i], pixels = pixels[:XSize], pixels[XSize:]
	}
	fmt.Println(picture2)
	fmt.Println()
}

func testFindMin() {
	fmt.Println(Min(10, 5, 99, 4, 2))
	fmt.Println()
}

func Min(a ...int) int {
	min := int(^uint(0) >> 1)
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func testAppend() {
	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	x = []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println()
}

const (
	C0 int = 1
	C1     = 2
	C2
	C3
	C4
	C5 = iota
	C6
	C7
	C8
	C9 = 111
	C10
)

func testConst() {
	fmt.Println(C0, C1, C2, C3, C4, C5, C6, C7, C8, C9, C10)
	fmt.Println()
}
