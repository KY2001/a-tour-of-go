package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const allowedError = 1e-10
	z := 1.0
	for {
		z_next := z - (z*z-x)/(2*z)
		if math.Abs(z-z_next) < allowedError {
			break
		} else {
			z = z_next
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(10))
}
