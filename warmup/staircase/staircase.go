// package main

// import (
// 	"fmt"
// )

//	func staircase(n int32) {
//		for i := int32(0); i < n; i++ {
//			row := ""
//			for j := int32(0); j < n-i-1; j++ {
//				row = fmt.Sprintf("%s ", row)
//			}
//			for j := int32(0); j <= i; j++ {
//				row = fmt.Sprintf("%s#", row)
//			}
//			fmt.Println(row)
//		}
//	}
//
//	func main() {
//		var n int32 = 8
//		staircase(n)
//	}
package main

import (
	"fmt"
)

func staircasePattern(n int) {
	for i := 0; i < n; i++ {
		r := ""
		for j := 0; j < n-i-1; j++ {
			r += " "
		}
		for j := 0; j <= i; j++ {
			r += "A"
		}
		fmt.Println(r)
	}
}

func main() {
	n := 5
	staircasePattern(n)
}
