package main

import (
	"fmt"
	"math"
	"time"
)

const (
	big   = 1 << 100
	small = big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x*10 + 1
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return lim
	} else {
		return v
	}
}

// 公式：z=z-(z*z-x)/2z
func sqrt(x float64) float64 {
	z := 1.0
	for {
		temp := z - (z*z-x)/(2*z)
		fmt.Println(temp)
		if temp == z || math.Abs(temp-z) < 0.000000000001 {
			break
		}
		z = temp
	}
	return z
}

func main() {
	defer fmt.Println("world")
	defer fmt.Println("1212312")
	fmt.Println("hello")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Friday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
