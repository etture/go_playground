package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	samoa := person{name: "Sean", age: 50}
	fmt.Println(samoa.name)

	sp := &samoa
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	s := person{age: 14}
	fmt.Println(s)
	fmt.Println(s.name)
}
