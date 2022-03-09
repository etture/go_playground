package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	b, err := ioutil.ReadFile("sample.lua")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	fmt.Println()
	fmt.Println(string(b))
	fmt.Println()
	fmt.Println(time.Now().UTC().Unix())
}
