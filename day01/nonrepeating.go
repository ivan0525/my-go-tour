package main

import (
	"fmt"
	"math"
)

func lengthOfNonRepeating(s string) int {
	start,maxLength:=0,0
	exist:=make(map[rune]int)
	for end,ch:=range []rune(s) {
		//if exist[ch]>=0 {
		//	// 如果map中存在当前出现的字符，更新start
		//	start= int(math.Max(float64(exist[ch]), float64(start)))
		//}
		v,ok:=exist[ch]
		if ok {
			start= int(math.Max(float64(v), float64(start)))
		}
		maxLength=int(math.Max(float64(maxLength), float64(end-start+1)))
		exist[ch]=end+1
	}
	return maxLength
}
func main() {
	fmt.Println(
		lengthOfNonRepeating("abcdefg"),
		lengthOfNonRepeating("abcabcs"),
		lengthOfNonRepeating("aas"),
		lengthOfNonRepeating("b"),
		lengthOfNonRepeating(" "),
	)
}
