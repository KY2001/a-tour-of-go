package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

var EPS float64 = 1e-10

func Sqrt(x float64) (float64, error) {
	z := 1.0
	for i := 0; i < 100; i++ {
		delta := (z*z - x) / (2 * z)
		if math.Abs(delta) < EPS {
			break
		}
		z -= delta
	}
	if x < 0 {
		return z, ErrNegativeSqrt(x)
	} else {
		return z, nil
	}
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
