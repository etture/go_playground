package main

import "fmt"

func main() {

	err := "hi"
	fmt.Println(err)

	if err := some(); err != "foo" {
		fmt.Println("haha1")
	}

	err := some()

}

func some() string {
	return "nil"
}
