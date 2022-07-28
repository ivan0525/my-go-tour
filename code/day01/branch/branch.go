package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	readFile1()
	readFile2()
	splitString()
	pointArray()
	// testArray()
}

func readFile1() {
	const filename = "aaa.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func readFile2() {
	const filename = "aaa.txt"
	// if语句中可以进行赋值，变量的作用域在if块内
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// fmt.Println(contents) // 访问不到
}

func splitString() {
	s := "hello"
	a := strings.Split(s, "he")
	b := s[3:]
	fmt.Println(b, a, s)
}

func testArray() {
	arr := [4]int{1, 2, 3, 4}
	// for循环遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d]=%d\n", i, arr[i])
	}
	fmt.Println("使用for...range进行遍历")
	for index, v := range arr {
		fmt.Printf("arr[%d]=%d\n", index, v)
	}
	fmt.Println(arr)
	arr1 := arr
	arr[1] = 100
	fmt.Println(arr == arr1)
	fmt.Println(arr1, arr)
}

func pointArray() {
	// arr := [3]*int{1, 2, 3}
	// fmt.Println(arr)
	arr := [...]int{1, 2}
	fmt.Println(arr)
	fmt.Println("struct数组")
	arr1 := [...]struct {
		name string
		age  int
	}{
		{"user1", 19},
		{"user2", 20},
	}
	fmt.Println(arr1)
}
