package main

import "fmt"

func main() {
	s := []int{1,2,3,4}
	sum:=0
	for _, num := range s {
		sum += num
	}
	fmt.Println(sum)

	for i, num:= range s {
		if num == 4 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "abc" {
		fmt.Println(i, c)
	}
}
