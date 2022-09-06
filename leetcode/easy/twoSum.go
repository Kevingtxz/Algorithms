package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
}

func twoSum(nums []int, target int) []int {
	myMap := map[int]int{}
	for i, num := range nums {
		if firstIdx, ok := myMap[target - num]; ok {
			return []int{firstIdx, i}
		}
		myMap[num] = i
	}
	return []int{};
}