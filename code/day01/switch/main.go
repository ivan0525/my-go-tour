package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
}
