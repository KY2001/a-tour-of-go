package main

import (
	"fmt"
	"math"
)

var EPS float64 = 1e-10

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 100; i++ {
		delta := (z*z - x) / (2 * z)
		if math.Abs(delta) < EPS {
			break
		}
		fmt.Println(i, delta)
		z -= delta
	}
	return z
}

func main() {
	fmt.Println(math.Abs(Sqrt(2) - math.Sqrt(2)))
}
