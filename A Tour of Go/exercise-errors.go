package main

import (
	"fmt"
	// "math"
)

type ErrNegativeSqrt struct {
	Target float64
	What   string
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%v: %v",
		e.What, e.Target)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &ErrNegativeSqrt{
			Target: x,
			What:   "cannot Sqrt negative number",
		}
	}
	z := float64(1)
	for i := 1; i < 11; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	if result, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	if result, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	// fmt.Println(Sqrt(2))
	// fmt.Println(Sqrt(-2))
}
