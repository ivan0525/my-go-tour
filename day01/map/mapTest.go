package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 7
	m[2] = 13
	m[0] = 131
	fmt.Println("map:", m)

	v1 := m[1]
	fmt.Println("v1: ", v1)

	// 调用内置的len方法，返回的是键值对的数目
	fmt.Println("len:", len(m))

	// 内置的delete可以从map中已出键值对
	delete(m, 1)
	fmt.Println("map:", m)
	key, prs := m[0]
	fmt.Println(key, prs)
}