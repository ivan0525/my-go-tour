package main

import "fmt"

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


func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
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
}
