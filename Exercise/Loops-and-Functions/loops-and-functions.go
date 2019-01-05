package main

import (
	"fmt"
	"math"
)

// Sqrt :: 平方根をニュートン法で計算
func Sqrt(x float64) float64 {
	const eps float64 = 1e-11

	z0 := float64(1)
	z1 := z0

	count := int(1)

	for ; ; count++ {
		z0 -= (z0*z0 - x) / (2 * z0)
		if math.Abs(z1-z0) < eps {
			break
		}
		z1 = z0
	}

	fmt.Printf("calcCount: %d\n", count)

	return z0
}

func main() {
	value := float64(2)

	fmt.Printf("calcSqrt : %.16f\n", Sqrt(value))
	fmt.Printf("mathSqrt : %.16f\n", math.Sqrt(value))
}
