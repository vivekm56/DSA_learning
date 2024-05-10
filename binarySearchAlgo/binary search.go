package main

import "fmt"

func BinarySearch(arr []int, value int) {
	min := 0
	max := len(arr)
	mid := 0
	for min < max {
		mid = (min + max) / 2
		if value == arr[mid] {
			break
		} else if value > arr[mid] {
			min = mid + 1
		} else {
			max = mid - 1
		}
		fmt.Println(min, max, mid)
	}
	if min >= max {
		// fmt.Println("item not found")
		if min != max {
			if arr[min] < value {
				fmt.Println("item not found, nearest less value =", arr[min])
			}
			if arr[max] < value {
				fmt.Println("item not found, nearest less value =", arr[max])
			}
			if arr[mid] < value {
				fmt.Println("item not found, nearest less value =", arr[mid])
			}
		} else {
			if arr[mid] < arr[min] || arr[mid] < arr[max] {
				fmt.Println("item not found, nearest less value =", arr[mid])
			} else {
				fmt.Println("item not found, nearest less value =", arr[min])
			}
		}
	} else {
		fmt.Println("item found at =", mid)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 15, 26, 37, 58, 89, 100}
	BinarySearch(arr, 5)
}
