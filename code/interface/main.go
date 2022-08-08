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

type SalaryCalculator interface {
	CalculatorSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculatorSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculatorSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculatorSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	name := MyString("Sam Anderson")
	fmt.Printf("name.FindVowels(): %c\n", name.FindVowels()) // name.FindVowels(): [a e o]
	pemp1 :=
		Permanent{
			empId:    1,
			basicpay: 5000,
			pf:       20,
		}
	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}
