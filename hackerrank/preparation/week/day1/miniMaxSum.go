package main

import (
	"fmt"
)

func main() {
	miniMaxSum([]int32{1,3,5,7,9})
}


func miniMaxSum(arr []int32) {
	var sum, maximum, minimun int64 = 0, int64(arr[0]), int64(arr[0])
	for _, x := range arr {
		sum += int64(x)
		if int64(x) > maximum {
			maximum = int64(x)
		} else if int64(x) < minimun {
			minimun = int64(x)
		}
	}
	fmt.Println(sum-maximum, sum-minimun)
}