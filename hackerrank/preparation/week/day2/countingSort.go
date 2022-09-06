package main

import (
	"fmt"
)

func main() {
	fmt.Println(countingSort([]int32{1,1,1,3,2}))
}

func countingSort(arr []int32) [100]int32 {
	var result [100]int32
	for _, i := range arr {
			result[i]++
	}
	return result
}