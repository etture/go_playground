package main

import "fmt"

type random struct {
	a int8
	b string
}

type testStruct struct {
	name string
	r    *random
}

func (r *random) String() string {
	if r == nil || r.a > 10{
		return fmt.Sprint("r is nil")
	}
	return fmt.Sprintf("random -> (a: %d, b: %s)", r.a, r.b)
}

type somee uint8

const (
	some1 somee = iota
	some2
	some3
)

func main() {
	t := testStruct{}
	fmt.Printf("struct: %#v\n", t)
	fmt.Println("is nil:", t.r == nil)

	r := random{a: 10, b: "hello"}
	r1 := random{a: 11, b: "hello"}
	var r2 *random = nil
	fmt.Println(r.String())
	fmt.Println(r1.String())
	fmt.Println(r2.String())

	fmt.Println("some1:", some1)
	fmt.Println("some2:", some2)
	fmt.Println("some3:", some3)

	var aa []random
	var aaa *[]random
	fmt.Println("aa:", aa)
	fmt.Println("aaa:", aaa)
	fmt.Println("aa==nil:", aa==nil)
	fmt.Println("aaa==nil:", aaa==nil)
	if aa == nil {
		fmt.Println("aa is nil~")
	}
	if aaa == nil {
		fmt.Println("aaa is nil~")
	}

	x := 0b00000101
	y := 0b00010001
	fmt.Printf("x: %08b, y: %08b\n", x, y)
	fmt.Printf("x&y: %08b, x|y: %08b\n", x&y, x|y)

	var z uint8 = 0b00000001
	var zz uint16 = uint16(z)
	fmt.Printf("(uint8) z << 7: %020b\n", z << 7)
	fmt.Printf("(uint8) z << 8: %020b\n", z << 8)
	fmt.Printf("(uint8) zz << 8: %020b\n", zz << 8)

	var randList []random = nil
	fmt.Println("randList:", randList)
	fmt.Println("randList==nil:", randList==nil)
	fmt.Println("len(randList):", len(randList))
}
