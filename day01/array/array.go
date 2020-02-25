package main

import "fmt"

func testPointArray(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 100
}

func main() {
	arr := [2]int{}
	fmt.Printf("arr: %p\n", &arr)
	testPointArray(arr)
	fmt.Println(len(arr))
}
