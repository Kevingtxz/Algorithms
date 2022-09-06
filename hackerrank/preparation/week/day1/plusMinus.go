package main

import (
	"fmt"
)

func main() {
	plusMinus([]int32{1,1,0,-1,-1})
}


func plusMinus(arr []int32) {
	var countPositives, countNegative, countZeros, length float32 = 0, 0, 0, float32(len(arr))
	for _, x := range arr {
		if x > 0 {
			countPositives++
		} else if x < 0 {
			countNegative++
		} else {
			countZeros++
		}
	}
	fmt.Println(countPositives/length)
	fmt.Println(countNegative/length)
	fmt.Println(countZeros/length)
}