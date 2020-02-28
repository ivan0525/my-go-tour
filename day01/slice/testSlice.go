package main

import "fmt"

func main() {
	//createSlice()
	//operateSlice()
	//overCap()
	copySlice()
}

// copy: 切片拷贝
func copySlice() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s1, s2)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
}

// 超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
func overCap() {
	data := [...]int{0, 1, 2, 3, 4, 10: 10}
	fmt.Println(&data[0])
	s := data[:2:3]
	fmt.Println(cap(s))
	fmt.Println(s)
	s = append(s, 100, 200)      // 一次 append 两个值，超出 s.cap 限制。
	fmt.Println(cap(s),&data[0], &s[0]) // 比对底层数组起始指针。
	//0xc000070060 0xc00001e0f0
}

// 操作切片
func operateSlice() {
	var s = [6]int{1, 2, 3, 4, 5, 6}
	ss := s[1:4]
	ss[0] = 99
	fmt.Println(s, ss)

	f := [4]struct {
		x int
	}{}
	g := f[:]
	f[0].x = 1
	fmt.Printf("%p, %p\n", &f, &f[0])
	fmt.Println(g)

	// append向slice尾部添加元素，返回新的slice
	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)
	var b = []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b)
	c := append(a, b...)
	fmt.Printf("slice c : %v\n", c)
	d := append(c, 7)
	fmt.Printf("slice d : %v\n", d)
	e := append(d, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e)
}

// 声明切片
func createSlice() {

	// 方法一
	var s []int
	// 方法二
	s1 := make([]int, 0)
	fmt.Println(s1)
	if s == nil {
		fmt.Println("是空的")
	}

	// 初始化赋值
	var s4 []int = make([]int, 4, 6)
	fmt.Println(s4)

	s5 := [6]int{1, 2, 3, 4, 5, 6}
	var s6 []int
	s6 = s5[1:]
	fmt.Println(s6)

}
