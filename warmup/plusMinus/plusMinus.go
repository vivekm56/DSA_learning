package main

import (
	"fmt"
	"strconv"
)

func plusMinus(arr []int32) {
	var n float64 = float64(len(arr))
	var negativeValueCount float64
	var positiveValueCount float64
	var zeroValueCount float64
	for _, v := range arr {
		if v < 0 {
			negativeValueCount++
		} else if v > 0 {
			positiveValueCount++
		} else {
			zeroValueCount++
		}
	}
	fmt.Println(strconv.FormatFloat(positiveValueCount/n, 'f', 5, 32))
	fmt.Println(strconv.FormatFloat(negativeValueCount/n, 'f', 5, 32))
	fmt.Println(strconv.FormatFloat(zeroValueCount/n, 'f', 5, 32))

}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
