package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

type Person struct {
	Name string
	Age  int
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	p := Person{"Bob", 18}
	fmt.Println(p)
	sp := &p
	fmt.Println(sp.Age) // 对结构体指针使用'.'，指针会自动解引用
	fmt.Println(Person{Name: "Sally"})
}
