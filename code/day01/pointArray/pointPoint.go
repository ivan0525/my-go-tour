package main

import "fmt"

func main() {
	arr := [5]int{}
	printArr(&arr)
	fmt.Println(arr)
	arr2 := [5]int{1, 2, 3, 4, 5}
	printArr(&arr2)
	fmt.Println(arr2)

	a := [3]int{1, 2, 3}
	var pa [3]*int

	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]
	}

	for i := 0; i < len(pa); i++ {
		fmt.Printf("pa[i]: %v\n", *pa[i])
	}
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Printf("arr[%d]=%d\n", i, v)
	}

}
