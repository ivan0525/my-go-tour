package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	//name := "张鑫qq"
	//hobby := "睡觉"
	// 字符串拼接
	//ss := name + hobby
	//fmt.Println(ss)
	//ss1 := fmt.Sprintf("%s%s", name, hobby)
	//fmt.Println(ss1)

	// 字符串分割
	//s1 := strings.Split(ss1, "")
	//fmt.Println(s1)

	// 包含
	//fmt.Println(strings.Contains(ss, "睡觉"))

	// for a, s := range ss {
	// 	fmt.Println(a)
	// 	fmt.Printf("%c\n", s)
	// }

	//m := "big"
	//bytem := []byte(m)
	//fmt.Printf("%T\n", bytem[0])
	//bytem[0] = 'p'
	//fmt.Println(string(bytem[0]))
	//
	//s2 := "白萝卜"
	//runeS2 := []rune(s2)
	//fmt.Println(runeS2)
	//homework()
	//var s = "hello, 我喜欢编程!"
	//// 获取字符数量
	//length:=utf8.RuneCountInString(s)
	//// len获取字节长度
	//fmt.Println(length,len(s))
	//for i,v:=range []rune(s) {
	//	//fmt.Println(i,v)
	//	fmt.Printf("(%d,%c)\n",i,v)
	//}
	//s := "颜色感染是一个有趣的游戏。"
	//bs := []byte(s)      // string -> []byte
	//s = string(bs)       // []byte -> string
	//rs := []rune(s)      // string -> []rune
	//s = string(rs)       // []rune -> string
	//rs = bytes.Runes(bs) // []byte -> []rune
	//bs = Runes2Bytes(rs) // []rune -> []byte
	var s = "babad"
	fmt.Println(s[1:4])
	longestPalindrome(s)
}

func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
		fmt.Println(n)
	}
	return bs
}

// 统计字符串中的中文个数
func homework() string {
	var str = "hello张鑫大屌哥"
	var count int
	runeStr := []rune(str)
	for _, s := range runeStr {
		if s > 255 {
			count++
		}
	}

	math.Max(1, 2)
	fmt.Println(count)
	return ""
}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	for index, _ := range s {
		var len1 = expandAroundCenter(s, index, index)
		var len2 = expandAroundCenter(s, index, index+1)
		var len = 0
		if len1>len2 {
			len=len1
		}else {
			len=len2
		}
		if len > end-start {
			start = index - (len-1)/2
			end = index + len/2
		}

	}
	fmt.Println(start,end,s[start:end+1])
	return s[start:end]
}

func expandAroundCenter(s string, left int, right int) int {
	l := left
	r := right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
