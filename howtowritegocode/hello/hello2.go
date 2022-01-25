package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"example/user/hello2/morestrings"
)

func main() {
	fmt.Println("Hello, world2.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello, World", "Hello, Go"))
}
