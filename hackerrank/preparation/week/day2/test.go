package main

import (
	"fmt"
)

func main() {
	fmt.Println(flippingMatrix([][]int32{{1,2,3},{4,5,6},{9,8,7}}))
}

func flippingMatrix(matrix [][]int32) int32 {
	var sum int32 = 0 
	length := len(matrix)
	for i1 := 0; i1 < length; i1++ {
		for i2 := 0; i2 < length; i2++ {
			fmt.Println(i1, i2)
			for _, arr := range matrix {
				fmt.Println(arr)
			}
			if matrix[i1][i2] < matrix[i1][length-1-i2] {
				for i3 := 0; i3 < length; i3++ {
					matrix[i1][i3], matrix[i1][length-1-i3] = matrix[i1][length-1-i3], matrix[i1][i3]
				}
			}
		}
	}

	for i := 0; i < length; i++ {
		sum += matrix[i][i]
	}


	return sum
}