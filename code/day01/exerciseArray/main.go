package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := [5]int{1, 3, 5, 8, 7}
	fmt.Println(sumArr(arr))
	// 产生随机数
	rand.Seed(time.Now().Unix())
	for i:=0;i<len(arr);i++ {
		arr[i]=rand.Intn(1000)
	}
	fmt.Println(arr)
	fmt.Println(sumArr(arr))
	twoSum(arr, 8)

}

func sumArr(arr [5]int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func twoSum(arr [5]int, target int) {
	for i := 0; i < len(arr); i++ {
		sub := target - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == sub {
				fmt.Printf("(%d,%d)", i, j)
			}
		}
	}
}

