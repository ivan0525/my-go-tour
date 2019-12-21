package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := float64(1)
	tempZ := float64(0)
	for math.Abs(z-tempZ) > 0.0000000001 {
		tempZ = z
		z = (z + x/z) / 2
	}
	return z
}

func main() {
	fmt.Println(sqrt(2), math.Sqrt(2))
}
