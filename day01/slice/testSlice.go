package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "hello"
	fmt.Println(s)
	s = append(s, "world")
	fmt.Println(s)
	fmt.Println("len: ", len(s))
	var s1 = []string{"a", "b", "c"}
	fmt.Println(s1[:2])
}