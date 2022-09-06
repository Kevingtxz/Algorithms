package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func main() {
	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", quickSort(slice))
}


func quickSort(arr []int) []int {
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