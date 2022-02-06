package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (ss byLength) Len() int {
	return len(ss)
}

func (ss byLength) Less(i, j int) bool {
	return len(ss[i]) < len(ss[j])
}

func (ss byLength) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
