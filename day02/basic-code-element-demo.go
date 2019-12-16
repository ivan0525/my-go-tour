package main

import (
	"math/rand"
)

const maxRand = 40

func getRandomNumber(numRands int) (int, int) {
	var a, b int
	for i := 0; i < maxRand; i++ {
		if rand.Intn(maxRand) < maxRand/2 {
			a++
		} else {
			b++
		}
	}
	return a, b
}

// func main() {
// 	num := 40
// 	x, y := getRandomNumber(num)
// 	print("Result：", x, " + ", y, " = ", num, " ? ")
// 	println(x+y == num)
// }
