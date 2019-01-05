package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt :: Sqrt 関数で引数が負数の場合の Error タイプ
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt :: 平方根をニュートン法で計算
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	const eps float64 = 1e-11

	z0 := float64(1)
	z1 := z0

	for {
		z0 -= (z0*z0 - x) / (2 * z0)
		if math.Abs(z1-z0) < eps {
			break
		}
		z1 = z0
	}

	return z0, nil
}

func printSqrt(x float64) {
	sqrt, error := Sqrt(x)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(sqrt)
	}
}

func main() {
	printSqrt(2)
	printSqrt(-2)
}
