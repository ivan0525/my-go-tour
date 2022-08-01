package main

import (
	"fmt"
	"strings"
)

func add(x int, y int) int {
	return x + y
}

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	fmt.Println(add(1, 2))
	sum(1, 2)
	sum(1, 2, 3)
	s := []int{1, 2, 3, 4}
	sum(s...)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	jpgFunc := makeSuffixFunc(".jpg")
	textFunc := makeSuffixFunc(".txt")

	v1 := jpgFunc("test")
	v2 := textFunc("test.txt")

	fmt.Printf("v1: %v\n", v1)
	fmt.Printf("v2: %v\n", v2)
}
