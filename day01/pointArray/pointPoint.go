package main

import "fmt"

func main() {
	arr := [5]int{}
	printArr(&arr)
	fmt.Println(arr)
	arr2 := [5]int{1, 2, 3, 4, 5}
	printArr(&arr2)
	fmt.Println(arr2)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Printf("arr[%d]=%d\n", i, v)
	}

}
