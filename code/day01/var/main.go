package main

import "fmt"

var (
	name          string
	age           int
	hasGirlFriend bool
)

func main() {
	name = "张鑫"
	age = 22
	hasGirlFriend = false
	// Go语言中的变量声明了就必须使用,不然就编译不过去
	fmt.Print(age)                // 在终端输出要打印的内容
	fmt.Println(hasGirlFriend)    // 打印完指定的内容后会在后面加一个换行符
	fmt.Printf("name:%s\n", name) // %s: 占位符,使用name这个变量的值去替换占位符

	// 类型推导
	var studentName = "哈哈"
	fmt.Println(studentName)

	// 简短变量声明(只能在函数中使用)
	_s1 := false
	fmt.Println(_s1)
}
