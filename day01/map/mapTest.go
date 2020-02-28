package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map的基本使用
func basicMap() {
	// map类型的变量初始值为nil，需要通过make()函数来初始化
	// make(map[KeyType]ValueType, [cap]),
	//其中：KeyType为键的类型，ValueType表示键对应的值的类型，cap表示map的容量
	m := make(map[string]int, 6)
	m["age"] = 19
	m["quantity"] = 100
	fmt.Println(m, m["age"])
	fmt.Printf("type of m: %T\n", m)

	// 支持在生命的时候添加元素
	a := map[string]string{
		"username": "zhangxin",
		"age":      "19",
	}
	fmt.Println(a)

	// 判断某个键是否存在
	// 如果键存在，ok为true，v为对应的键值；否则，ok为false，v为对应类型的零值。
	v, ok := a["s"]
	fmt.Println(v) // ""
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("无对应键")
	}

	// map的遍历，使用for...range进行遍历
	// 第一个参数为键，第二个参数为键对应的值
	// 只想要遍历key的时候，可以不写第二个参数
	// 注意： 遍历map时的元素顺序与添加键值对的顺序无关。
	for k, v := range a {
		fmt.Println(k, v)
	}

	// 删除map中的键
	// 使用内建的delete()函数
	delete(a, "age")
	fmt.Println(a)
}

// 按照指定顺序对map进行遍历
func showByOrder() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
	m := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		m[key] = rand.Intn(100)
	}

	// 将map中的所有key存入切片keys
	keys := make([]string, 0, 200)
	for k := range m {
		keys = append(keys, k)
	}
	// 对切片进行排序
	sort.Strings(keys)
	fmt.Println(keys)
	for _, key := range keys {
		fmt.Println(key, m[key])
	}
}

// 值为map类型的切片
func mapSlice() {
	var sliceMap = make([]map[string]string, 4)
	for i, v := range sliceMap {
		fmt.Printf("index: %d value: %v\n", i, v)
	}
	// 对切片中的map元素进行初始化
	for i := range sliceMap {
		sliceMap[i] = make(map[string]string, 5)
	}
	sliceMap[0]["name"] = "小虎"
	sliceMap[1]["name"] = "zhangxin"
	fmt.Println(sliceMap)
}

// 值为切片类型的map
func sliceMap() {
	var sliceMap = make(map[string][]string, 3)
	//fmt.Println(sliceMap)
	key := "username"
	v, ok := sliceMap["username"]
	if ok {
		fmt.Println(v)
	} else {
		v = make([]string, 0, 2)
	}
	v = append(v, "北京", "上海")
	sliceMap[key] = v
	fmt.Println(cap(sliceMap[key]))
}

func main() {
	//basicMap()
	//showByOrder()
	//mapSlice()
	sliceMap()
}
