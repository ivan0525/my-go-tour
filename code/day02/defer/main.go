package main

import "fmt"

/*
1. 关键字defer用于注册延迟调用
2. 这些调用直到return前才被执行。因此，可以用来做资源清理。
3. 多个defer语句，按先进后出的方式执行。
4. defer语句中的变量，在defer声明时就决定了。
*/

func main() {
	defer test()
	defer fmt.Println("world")

	fmt.Println("hello")
}

func test() {
	a := 1
	fmt.Println(a)
}
