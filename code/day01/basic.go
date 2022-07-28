package main

import "fmt"

var aa = 1
var ss = false

func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("a=%d\nb=%q\n", a, b)
}

func variableInitialValue() {
	var (
		a, b int    = 10, 10
		s    string = "hello"
	)
	fmt.Println(a, b, s)
}

func enums() {
	const (
		python = iota
		_
		java
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(python, java, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}



func main() {
	//variableZeroValue()
	//variableInitialValue()
	//
	//enums()
}
