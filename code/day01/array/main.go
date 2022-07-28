package main

import (
	"fmt"
)

func main() {
	var s [5]int
	s[4] = 100
	fmt.Println(s)
	fmt.Println("set:", s)
	fmt.Println("get:", s[4])
	fmt.Println("len", len(s))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
}
