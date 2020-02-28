package main

import (
	"fmt"
)

func test() {
	a := 10
	b := &a
	fmt.Printf("a:%d type:%T\n", a, a)
	fmt.Printf("&a:%p type:%T\n", &a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b)
}

// 空指针的判断
// 结论：当一个指针定义后没有被分配到任何变量时，它的值为nil
func emptyPoint() {
	var p *string
	fmt.Printf("p: %T\np的值是%v\n", p, p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}

// 引用类型的变量，在使用的时候不仅要声明它，还需要为其分配内存空间。
// 对于值类型的声明不需要主动分配内存空间，因为它们在声明的时候已经分配好了内存空间。
func assignMemory() {
	var a *int
	*a = 1
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
}

// 使用new和make内建的方法来为引用类型变量分配内存空间
// func new(Type) *Type
// 1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
// 2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
// new用于类型的内存分配，并且内存对应的值为该类型的零值，返回的是指向类型的指针
func useNew() {
	a := new(int) // 返回的是一个指向int类型内存地址的指针
	*a = 100
	fmt.Println(*a)
}

// 使用make来进行内存分配
// func make(t Type, size ...IntegerType) Type
// 它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，
//而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
func useMake() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["test"] = 100
	fmt.Println(b)
}

// 练习
func exercise() {
	a := new(int)
	fmt.Println(a)
	b := new(int)
	b = a
	*b = 20
	fmt.Println(*a, *b)
	arr := new([]int)
	fmt.Printf("%T\n", *arr)
	*arr = []int{1, 2, 3}
	fmt.Println(*arr)
}

// 使用牛顿法求平方根
func binarySqrt(n int) int {
	x := n
	for x*x > n {
		x = (x + n/x) / 2
	}
	return x
}

func testMore(a *[]int) {
	*a = append(*a, *a...)
	a = nil
	//fmt.Println(*a)
}

// 指针比较
func pointerCompare() {
	type Myint int64
	type Ta *int64
	type Tb *Myint

	var pa0 Ta
	var pa1 *int64
	var pb0 Tb
	var pb1 *Myint
	fmt.Println(pa0 == pa1)
	fmt.Println(pb1 == pb0)
	pa0 = pa1
}

func numberOfSteps(num int) int {
	count := 0
	for num != 0 {
		// 偶数
		if num/2 == 0 {
			num = num / 2
		} else {
			num--
		}
		count++
	}
	return count
}

func main() {
	//test()
	//emptyPoint()
	//assignMemory()
	//useNew()
	//useMake()
	//exercise()
	//fmt.Println(binarySqrt(3))
	//s:=[]int{1,2}
	//fmt.Println(s)
	//testMore(&s)
	//fmt.Println(s)
	//pointerCompare()
	numberOfSteps(14)
}
