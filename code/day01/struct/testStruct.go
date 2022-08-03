package main

import "fmt"

type person struct {
	username, city string
	age            int
	sex            byte
	subject        []string
}

// 结构体实例化
func example() {
	var person1 person
	person1.age = 19
	person1.city = "Guangdong"
	person1.sex = 1
	person1.subject = []string{"math", "chinese"}
	fmt.Println(person1)
	fmt.Printf("type of person1: %T\n", person1)
	fmt.Printf("%#v\n", person1)
}

// 匿名结构体
func nonStruct() {
	var person struct {
		name     string
		age, sex int8
	}
	person.name = "张三"
	person.age = 20
	person.sex = 1
	fmt.Printf("%#v\n", person)
}

// 结构体内存布局
func structMemory() {
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}

type student struct {
	name string
	age  int
}

func question() {
	// m的值指向同一个结构体指针*student
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}
	for _, stu := range stus {
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, v)
		//fmt.Println(k, "=>", v.name)
	}
}

type animal struct {
	name     string
	age, sex int8
}

// 自定义结构体构造函数
func newAnimal(name string, age, sex int8) *animal {
	return &animal{
		name: name,
		age:  age,
		sex:  sex,
	}
}

// 结构体字面量形式
func literalStruct() {
	type Book struct {
		page int
	}

	//var book = Book{}
	//p:=&book.page
	//*p=10
	//fmt.Printf("%v\n",book)
	//var a = 1
	//fmt.Printf("%T",&a)

	//fmt.Println(b)
	//var s = "hello"
	//fmt.Println()
	a := 1
	b := &a
	fmt.Println(b)
	(&Book{}).page = 1
}

func swap1(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func swap2(a, b int) (int, int) {
	return b, a
}

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(val int) {
	if node == nil {
		return
	}
	node.value = val
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

// 树的遍历
// 先序遍历
func (node *treeNode) preOrder() {
	if node == nil {
		return
	}
	node.print()
	node.left.preOrder()
	node.right.preOrder()
}

// 中序遍历
func (node *treeNode) midOrder() {
	if node == nil {
		return
	}

	node.left.midOrder()

	node.print()
	node.right.midOrder()
}

// 后序遍历
func (node *treeNode) lastOrder() {
	if node == nil {
		return
	}

	node.left.lastOrder()
	node.right.lastOrder()
	node.print()
}

func makeTree() {
	root := treeNode{value: 1}
	root.left = &treeNode{}
	root.right = createNode(4)
	root.left.right = new(treeNode)
	root.right.left = &treeNode{}
	//root.print()
	//root.setValue(4)
	//root.print()
	//var pRoot *treeNode
	//fmt.Println(pRoot)
	//pRoot.setValue(100)
	//pRoot=&root
	//pRoot.setValue(500)
	//pRoot.print()
	root.preOrder()
}

type Point struct {
	x int
	y int
}

// func main() {
// 	//makeTree()
// 	//example()
// 	//nonStruct()
// 	structMemory()
// 	//question()
// 	//var dog = newAnimal("小虎", 4, 0)
// 	//var cat = newAnimal("小花", 3, 1)
// 	//fmt.Println(dog, cat)
// 	//literalStruct()

// 	//x, y := 1, 2
// 	//swap1(&x, &y)
// 	//fmt.Println(x, y)
// 	//fmt.Println(swap2(x,y))
// 	p := Point{1, 2}
// 	p1 := &p
// 	(*p1).x = 1e9
// 	fmt.Println(p)
// 	v := Point{y: 1}
// 	v1 := Point{}
// 	v2 := &Point{3, 4}
// 	fmt.Println(v, v1, v2)
// }
