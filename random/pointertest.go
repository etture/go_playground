package main

import "fmt"

func main() {
	//passByValue()
	//passByPointer()

	var ps *PageSize = func() *PageSize { p := PageSize(15); return &p}()

	fmt.Println(ps)
}

type person struct {
	name string
}

func passByValue() {
	p := person{name: "Richard"}
	renameByValue(p)
	fmt.Println(p)
}

func renameByValue(p person) {
	p.name = "test"
}

func passByPointer() {
	p := person{name: "Richard"}
	renameByPointer(&p)
	fmt.Println(p)
}

func renameByPointer(p *person) {
	p.name = "test"
}

type PageSize uint32


