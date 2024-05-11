package main

import (
	"fmt"
	"math"
)

func powerSum(X int32, N int32, numb int32) int32 {
	// Write your code here
	var pwr = int32(math.Pow(float64(numb), float64(N)))

	if pwr < X {
		return powerSum(X, N, numb+1) + powerSum(X-pwr, N, numb+1)
	}

	if pwr == X {
		return 1
	} else {
		return 0
	}
}
func main() {
	var X int32 = 100
	var N int32 = 3
	result := powerSum(X, N, 1)
	fmt.Println(result)
}
