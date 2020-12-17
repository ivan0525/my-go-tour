package main

import (
	"fmt"
	"math"
)
// 牛顿开平方公式：z=z-(z*z-x)/(2*z)
func sqrt (x float64) float64 {
	z:=1.0
	for {
		temp:=z-(z*z-x)/(2*z)
		if temp == z || math.Abs(temp - z) < 0.000000000001 {
			break
		}
		z = temp
	}
	return z
}
func main() {
	fmt.Println(sqrt(2.0), math.Sqrt(2.0))
}
