package main

import "fmt"

func main() {
	defer test()
	defer fmt.Println("world")

	fmt.Println("hello")
}

func test() {
	a := 1
	fmt.Println(a)
}
