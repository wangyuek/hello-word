package main

import (
	"fmt"
)

func my_sqrt(x float32) float32 {
	var z float32 = 1.0
	for i := 1; i <= 100; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("i=%d,x=%f,z=%f\n", i, x, z)
		z = z + 1.0
	}
	return 0.0
}

func main() {
	var x float32
	x = 100
	my_sqrt(x)
}
