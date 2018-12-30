package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	const eps float64 = 1e-11

	z0 := float64(1)
	z1 := z0

	loopCount := int(0)

	for {
		z0 -= (z0*z0 - x) / (2 * z0)
		if math.Abs(z1-z0) < eps {
			break
		}
		z1 = z0
		loopCount++
	}

	fmt.Printf("loopCount: %d\n", loopCount)

	return z0
}

func main() {
	value := float64(2)

	fmt.Printf("calcSqrt : %.16f\n", sqrt(value))
	fmt.Printf("mathSqrt : %.16f\n", math.Sqrt(value))
}
