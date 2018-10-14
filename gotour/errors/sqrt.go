package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cant Sqrt negative number: %v", float64(err))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x
	for z*z-x > 0 {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%f \n", z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(-2))
}
