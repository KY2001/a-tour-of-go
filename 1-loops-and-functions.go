package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const allowedError = 1e-10
	z := 1.0
	for {
		delta := (z*z - x) / (2 * z)
		if math.Abs(delta) < allowedError {
			break
		} else {
			z -= delta
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(10))
}
