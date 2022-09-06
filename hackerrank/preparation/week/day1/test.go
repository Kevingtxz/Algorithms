package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedian([]int32{1,473,35,231,51,642,6442,3,142,53,566,234,321,314,52,623,43}))
}

func findMedian(arr []int32) int32 {
	quickSort(arr)
	return arr[len(arr)/2]
}

func quickSort(arr []int32) []int32 {
	if len(arr) < 2 {
			return arr
	}
	 
	left, right, middle := 0, len(arr)-1, len(arr)/2
	 
	var pivot int
	if arr[left] > arr[right] && arr[left] < arr[middle] || arr[left] < arr[right] && arr[left] > arr[middle] {
		pivot = left
	}	else if arr[middle] > arr[right] && arr[middle] < arr[left] || arr[middle] < arr[right] && arr[middle] > arr[left] {
		pivot = middle
	}	else {
		pivot = right
	}	 
	 
	arr[pivot], arr[right] = arr[right], arr[pivot]
	 
	for i, _ := range arr {
			if arr[i] < arr[right] {
					arr[left], arr[i] = arr[i], arr[left]
					left++
			}
	}
	 
	arr[left], arr[right] = arr[right], arr[left]
	 
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	 
	return arr
}