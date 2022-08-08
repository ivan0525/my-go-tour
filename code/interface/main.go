package main

import "fmt"

// 定义interface
type VowelsFinder interface {
	FindVowels() []rune
}

// 定义类型MyString
type MyString string

// 给receiver MyString添加FindVowels() []rune方法
// 这样MyString就implement接口VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Sam Anderson")
	fmt.Printf("name.FindVowels(): %c\n", name.FindVowels()) // name.FindVowels(): [a e o]
}
