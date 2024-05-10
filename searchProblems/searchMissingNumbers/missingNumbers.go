package main

import (
	"fmt"
	"sort"
)

func missingNumbers(arr []int32, brr []int32) []int32 {
	// Write your code here
	m1 := make(map[int32]int32)
	m2 := make(map[int32]int32)
	result := []int32{}
	for _, v := range arr {
		m1[v] = m1[v] + 1
	}
	for _, v := range brr {
		m2[v] = m2[v] + 1
	}
	fmt.Println(m1, m2)
	for k, v := range m1 {
		if _, ok := m2[k]; !ok {
			result = append(result, k)
			// fmt.Println(k)
		}
		if val, ok := m2[k]; ok && val != v {
			result = append(result, k)
			// fmt.Println(k)
		}
	}
	for k, _ := range m2 {
		if _, ok := m1[k]; !ok {
			result = append(result, k)
			// fmt.Println(k)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func main() {
	arr := []int32{203, 204, 205, 206, 207, 208, 203, 204, 205, 206}
	brr := []int32{203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204, 222}
	finaleResult := missingNumbers(arr, brr)
	fmt.Println(finaleResult)
}
