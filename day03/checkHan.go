package main

import (
	"fmt"
	"unicode"
)

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	i, cur, pre, res := 2, 2, 1, 0
	for i < n{
		i++
		res = (cur+pre)/1000000007
		pre = cur
		cur = res
	}
	return res
}

func main() {
	var str = "Hello, 张鑫"
	for _, s := range str {
		if unicode.Is(unicode.Han, s) {
			fmt.Printf("%v\n", s)
		}
	}
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	mapRoman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	num := 0
	next := 0 // 后一字符代表的值
	for i := len(s); i > 0; i-- {
		val := mapRoman[s[i-1:i]] // 当前值
		if val < next {
			num -= val
		} else {
			num += val
		}
		next = val
	}
	return num
}
