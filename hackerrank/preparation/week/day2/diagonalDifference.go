package main

import (
	"fmt"
)

func main() {
	fmt.Println(diagonalDifference([][]int32{{1,2,3},{4,5,6},{9,8,9}}))
}

func diagonalDifference(arr [][]int32) int32 {
	var result int32
	for i := 0; i < len(arr); i++ {
		result = result + arr[i][i] - arr[i][len(arr)-1-i]
	}
	if result < 0 {
		return -result
	}
	return result
}
