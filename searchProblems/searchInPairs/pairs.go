package main

import "fmt"

func pairs(k int32, arr []int32) int32 {
	var count int32 = 0
	exist := make(map[int32]bool)
	for _, number := range arr {
		exist[number] = true
		if _, ok := exist[number+k]; ok {
			count++
		}
		if _, ok := exist[number-k]; ok {
			count++
		}
	}
	return count
}

func main() {
	arr := []int32{1, 3, 5, 8, 6, 4, 2}
	var k int32 = 2
	a := pairs(k, arr)
	fmt.Println(a)
}
