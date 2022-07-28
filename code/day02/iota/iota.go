package main

import "fmt"

func main() {
	const (
		Filed = iota - 1
		UnKnown
		Succeeded
	)
	fmt.Print(Filed, UnKnown, Succeeded)
}
