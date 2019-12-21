package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	x, y := swap(1, 2)
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
