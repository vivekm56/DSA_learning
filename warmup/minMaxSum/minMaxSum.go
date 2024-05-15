package main

import (
	"fmt"
)

func minMaxSum(arr []int32) {

	min, max := arr[0], arr[0]
	var sum int
	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += int(num)
	}
	fmt.Printf("%d %d\n", sum-int(max), sum-int(min))

}
func main() {
	arr := []int32{1, 2, 3, 4, 5}
	minMaxSum(arr)
}
