package main

import (
	"fmt"
)

func main() {
	fmt.Println(lonelyInteger([]int32{1,1,0,-1,-1}))
}


func lonelyInteger(arr []int32) int32 {
	var myMap = map[int32]bool{}
	for i := 1; i < 10; i++ {
		myMap[int32(i)] = false
	}
	for _, x := range arr {
		myMap[x] = !myMap[x]
	}
	for x := range myMap {
		if myMap[x] == true {
			return x
		}
	}
	return -1
}