package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "张鑫qq"
	hobby := "睡觉"
	// 字符串拼接
	ss := name + hobby
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, hobby)
	fmt.Println(ss1)

	// 字符串分割
	s1 := strings.Split(ss1, "")
	fmt.Println(s1)

	// 包含
	fmt.Println(strings.Contains(ss, "睡觉"))

	// for a, s := range ss {
	// 	fmt.Println(a)
	// 	fmt.Printf("%c\n", s)
	// }

	m := "big"
	bytem := []byte(m)
	fmt.Printf("%T\n", bytem[0])
	bytem[0] = 'p'
	fmt.Println(string(bytem[0]))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	fmt.Println(runeS2)
	homework()
}

// 统计字符串中的中文个数
func homework() {
	var str = "hello张鑫大屌哥"
	var count int
	runeStr := []rune(str)
	for _, s := range runeStr {
		if s > 255 {
			count++
		}
	}
	fmt.Println(count)
}
