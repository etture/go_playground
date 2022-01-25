package main

import "fmt"

func main() {
	passByValue()
	passByPointer()
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

