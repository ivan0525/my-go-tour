package main

// 每个go程序都是由包组成的
// 程序运行的入口包是main
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("The time is:", time.Now())
}
